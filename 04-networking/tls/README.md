# crypto/tls

## Overview
`crypto/tls` implements TLS, the protocol behind HTTPS. Most of the time you
use it indirectly: `http.Client` and `http.Server` handle TLS for you, and you
only supply a `tls.Config` to set a minimum version, trust roots, client
certificates or a custom certificate. Go negotiates TLS 1.3 by default and has
safe defaults, so the best configuration is usually a small one.

## Core API

| Function / Type | Purpose |
|---|---|
| `tls.Config` | Settings: `Certificates`, `RootCAs`, `ServerName`, `MinVersion`, `ClientAuth`, `ClientCAs` |
| `tls.Dial`, `tls.Listen` | TLS over TCP |
| `tls.Conn.ConnectionState()` | Negotiated version, cipher suite, peer certificates |
| `http.Transport{TLSClientConfig: ...}` | TLS settings for an HTTP client |
| `http.Server{TLSConfig: ...}`, `ServeTLS`, `ListenAndServeTLS` | HTTPS server |
| `tls.LoadX509KeyPair` | Load a certificate and key from files |
| `x509.NewCertPool`, `AddCert`, `AppendCertsFromPEM` | Custom trust roots |
| `tls.VersionName`, `tls.CipherSuiteName` | Readable names |
| `httptest.NewTLSServer` | HTTPS test server |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/tls/
```

## Gotchas
- **Never set `InsecureSkipVerify: true` outside tests.** It turns off certificate checking and makes TLS pointless. If you need to trust a private CA, put it in `RootCAs`.
- **Set `MinVersion: tls.VersionTLS12`** (or 1.3) explicitly in servers and clients you configure. Leave cipher suites alone; the defaults are curated.
- **The server name must match the certificate.** Use a hostname that is in the certificate's DNS or IP SANs; the Common Name alone is ignored.
- **`Certificates` is for the side that proves its identity** (servers, and clients using mTLS). **`RootCAs` is for the side that verifies** (clients). **`ClientCAs` + `ClientAuth` verify clients** on the server.
- **The handshake happens on the first read or write** (or explicit `Handshake()`), so certificate errors appear there, not at `Dial`/`Accept` of a plain `net.Conn`. `tls.Dial` does the handshake for you.
- **Reuse `http.Transport` and clients;** each new transport re-handshakes and defeats connection pooling.
- **Certificate errors are typed:** use `errors.As` with `*tls.CertificateVerificationError`, `x509.HostnameError` or `x509.UnknownAuthorityError`.
- **Reload certificates without restarting** with `tls.Config.GetCertificate`, which is called per handshake.
- **Do not generate certificates in production code.** The self-signed helper in these examples is for tests only.

## Java comparison

| Java | Go |
|---|---|
| `SSLContext`, `SSLSocket` | `tls.Config`, `tls.Conn` |
| `KeyStore` / `TrustStore` | `tls.Certificate` / `x509.CertPool` |
| `HttpsURLConnection` hostname verifier | `ServerName` + built-in verification |
| Spring Boot `server.ssl.*` | `http.Server{TLSConfig}` + `ServeTLS` |
| mutual TLS (`clientAuth=need`) | `ClientAuth: tls.RequireAndVerifyClientCert` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/crypto/tls
- Source: https://cs.opensource.google/go/go/+/master:src/crypto/tls/
