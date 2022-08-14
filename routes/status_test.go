package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GorginZ/georgia-hello-world/metadata"
)

func Test_StatusBody(t *testing.T) {
	tests := map[string]struct {
		wantBody string
		request  *http.Request
		w        httptest.ResponseRecorder
	}{
		"status-should-be-valid-json": {
			wantBody: `{"my-application":{"version":"1.0","description":"text","sha":"abc53458585"}}\n`,
			request:  httptest.NewRequest("GET", "/status", nil),
			w:        *httptest.NewRecorder(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// overide our baked in values these don't change at runtime just testing shape
			metadata.Version = "1.0"
			metadata.Sha = "abc53458585"
			metadata.Description = "text"

			HandleStatus(&tt.w, tt.request)
			response := tt.w.Result()
			response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf(err.Error())
			}

			// this is awful, but the \n in the string literal is being escaped, so I need to put the newline char 'back'
			want := strings.Replace(tt.wantBody, `\n`, "\n", -1)
			if string(body) != string(want) {
				t.Errorf("got %q, want %q", string(body), tt.wantBody)
			}
		})
	}
}
