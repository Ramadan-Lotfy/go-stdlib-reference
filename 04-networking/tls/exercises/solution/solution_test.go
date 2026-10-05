package solution

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewHTTPSClient(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	}))
	defer srv.Close()

	pool := x509.NewCertPool()
	pool.AddCert(srv.Certificate())
	client := NewHTTPSClient(pool)

	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("trusted server should work: %v", err)
	}
	resp.Body.Close()

	// A client with an empty pool must reject the same server.
	if resp, err := NewHTTPSClient(x509.NewCertPool()).Get(srv.URL); err == nil {
		resp.Body.Close()
		t.Error("untrusted server should be rejected")
	}
}

func TestClientSettings(t *testing.T) {
	client := NewHTTPSClient(x509.NewCertPool())
	if client.Timeout != 5*time.Second {
		t.Errorf("Timeout = %v, want 5s", client.Timeout)
	}
	cfg := client.Transport.(*http.Transport).TLSClientConfig
	if cfg.MinVersion < tls.VersionTLS12 {
		t.Errorf("MinVersion = %x, want at least TLS 1.2", cfg.MinVersion)
	}
	if cfg.InsecureSkipVerify {
		t.Error("InsecureSkipVerify must be false")
	}
}
