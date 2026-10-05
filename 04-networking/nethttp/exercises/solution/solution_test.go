package solution

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func do(h http.Handler, method, target, key string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, nil)
	if key != "" {
		req.Header.Set("X-API-Key", key)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestAPI(t *testing.T) {
	h := NewAPI()
	tests := []struct {
		name         string
		method, path string
		key          string
		wantCode     int
		wantBody     string
	}{
		{"health", "GET", "/health", "", 200, "ok"},
		{"no key", "GET", "/users/5", "", 401, "unauthorized"},
		{"wrong key", "GET", "/users/5", "nope", 401, "unauthorized"},
		{"found", "GET", "/users/5", "secret", 200, `{"id":5,"name":"user5"}`},
		{"not a number", "GET", "/users/abc", "secret", 400, "invalid id"},
		{"unknown user", "GET", "/users/500", "secret", 404, "user not found"},
		{"wrong method", "POST", "/users/5", "secret", 405, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := do(h, tt.method, tt.path, tt.key)
			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantCode)
			}
			if got := strings.TrimSpace(rec.Body.String()); tt.wantBody != "" && got != tt.wantBody {
				t.Errorf("body = %q, want %q", got, tt.wantBody)
			}
		})
	}
}
