# crypto/tls exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/tls/exercises/...
```

## Exercise: NewHTTPSClient

Write `NewHTTPSClient(pool *x509.CertPool) *http.Client` that:

- trusts only the certificates in `pool` (`RootCAs`)
- requires at least TLS 1.2
- never sets `InsecureSkipVerify`
- has a 5 second timeout

Test it with `httptest.NewTLSServer`: add `srv.Certificate()` to a pool, make a
request, and check that a client built from an **empty** pool rejects the same
server.

Hints: `http.Client`, `http.Transport{TLSClientConfig: &tls.Config{...}}`,
`x509.NewCertPool`.

## Bonus
Add mutual TLS: write `NewMTLSServerConfig(cert tls.Certificate, clientCAs *x509.CertPool) *tls.Config`
that requires and verifies client certificates (`tls.RequireAndVerifyClientCert`).
