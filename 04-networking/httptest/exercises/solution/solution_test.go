package solution

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.Write([]byte(`{"name":"ada"}`))
		case "/bad":
			w.Write([]byte(`not json`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	var v struct{ Name string }
	if err := GetJSON(context.Background(), srv.Client(), srv.URL+"/ok", &v); err != nil || v.Name != "ada" {
		t.Errorf("ok case: %+v, %v", v, err)
	}

	err := GetJSON(context.Background(), srv.Client(), srv.URL+"/missing", &v)
	if err == nil || !strings.Contains(err.Error(), "404") {
		t.Errorf("404 case: err = %v", err)
	}

	if err := GetJSON(context.Background(), srv.Client(), srv.URL+"/bad", &v); err == nil {
		t.Error("invalid JSON should fail")
	}
}

type failingTransport struct{}

var errBoom = errors.New("boom")

func (failingTransport) RoundTrip(*http.Request) (*http.Response, error) { return nil, errBoom }

func TestGetJSONTransportError(t *testing.T) {
	client := &http.Client{Transport: failingTransport{}}
	var v any
	err := GetJSON(context.Background(), client, "http://example.invalid", &v)
	if !errors.Is(err, errBoom) {
		t.Errorf("err = %v, want it to wrap errBoom", err)
	}
}
