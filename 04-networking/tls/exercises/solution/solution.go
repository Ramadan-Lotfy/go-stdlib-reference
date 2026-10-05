// Package solution holds the reference solutions for the crypto/tls exercises.
package solution

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"
)

// NewHTTPSClient returns a client that trusts only the certificates in pool,
// requires at least TLS 1.2, never skips verification and gives up after
// five seconds.
func NewHTTPSClient(pool *x509.CertPool) *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:    pool,
				MinVersion: tls.VersionTLS12,
			},
		},
	}
}
