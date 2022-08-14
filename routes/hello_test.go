package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RootStatus(t *testing.T) {
	tests := map[string]struct {
		wantCode int
		request  *http.Request
		w        httptest.ResponseRecorder
	}{
		"get-root-should-return-200": {
			wantCode: 200,
			request:  httptest.NewRequest("GET", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"should-deny-POST": {
			wantCode: 405,
			request:  httptest.NewRequest("POST", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"should-deny-DELETE": {
			wantCode: 405,
			request:  httptest.NewRequest("DELETE", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"should-deny-PUT": {
			wantCode: 405,
			request:  httptest.NewRequest("PUT", "/", nil),
			w:        *httptest.NewRecorder(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			HandleRoot(&tt.w, tt.request)
			response := tt.w.Result()
			if response.StatusCode != tt.wantCode {
				t.Errorf("got %v, want %v", response.Status, tt.wantCode)
			}
		})
	}
}

func Test_RootBody(t *testing.T) {
	tests := map[string]struct {
		wantBody string
		request  *http.Request
		w        httptest.ResponseRecorder
	}{
		"happy-hit-root-should-be-valid-json": {
			wantBody: `{"message":"Hello World!"}`,
			request:  httptest.NewRequest("GET", "/", nil),
			w:        *httptest.NewRecorder(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			HandleRoot(&tt.w, tt.request)
			response := tt.w.Result()
			response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf(err.Error())
			}
			if string(body) != string(tt.wantBody) {
				t.Errorf("got %q, want %q", string(body), tt.wantBody)
			}
		})
	}
}
