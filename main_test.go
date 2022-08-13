package main

import (
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
		"getRootShouldReturn200": {
			wantCode: 200,
			request:  httptest.NewRequest("GET", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"denyPOST": {
			wantCode: 405,
			request:  httptest.NewRequest("POST", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"denyDELETE": {
			wantCode: 405,
			request:  httptest.NewRequest("DELETE", "/", nil),
			w:        *httptest.NewRecorder(),
		},
		"denyPUT": {
			wantCode: 405,
			request:  httptest.NewRequest("PUT", "/", nil),
			w:        *httptest.NewRecorder(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			handleRoot(&tt.w, tt.request)
			response := tt.w.Result()
			if response.StatusCode != tt.wantCode {
				t.Errorf("got %v, want %v", response.Status, tt.wantCode)
			}
		})
	}
}
