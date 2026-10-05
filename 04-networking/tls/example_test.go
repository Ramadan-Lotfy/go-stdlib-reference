package tlsref_test

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

// selfSigned creates a throwaway certificate for 127.0.0.1 and localhost,
// plus a pool that trusts it. Real services load certificates from files or
// a certificate manager instead.
func selfSigned() (tls.Certificate, *x509.CertPool) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true, // lets the certificate act as its own trust root
		DNSNames:              []string{"localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	parsed, err := x509.ParseCertificate(der)
	if err != nil {
		panic(err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(parsed)
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}, pool
}

// Example_httptestTLS runs an HTTPS server with a throwaway certificate and
// uses the client that trusts it.
func Example_httptestTLS() {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "secure")
	}))
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL) // srv.Client trusts the test certificate
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(strings.HasPrefix(srv.URL, "https://"), string(body))
	fmt.Println(resp.TLS.HandshakeComplete, tls.VersionName(resp.TLS.Version))
	// Output:
	// true secure
	// true TLS 1.3
}

// Example_untrusted shows what an unknown certificate looks like to a client.
func Example_untrusted() {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()

	_, err := http.Get(srv.URL) // the default client does not trust the test CA
	var cve *tls.CertificateVerificationError
	fmt.Println(errors.As(err, &cve))

	var unknown x509.UnknownAuthorityError
	fmt.Println(errors.As(err, &unknown))
	// Output:
	// true
	// true
}

// Example_selfSigned builds a TLS server and client from scratch and checks
// the negotiated connection.
func Example_selfSigned() {
	cert, pool := selfSigned()
	ln, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		io.WriteString(conn, "secure hello\n") // the first write triggers the handshake
	}()

	conn, err := tls.Dial("tcp", ln.Addr().String(), &tls.Config{
		RootCAs:    pool,        // who we trust
		ServerName: "localhost", // name that must be in the certificate
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	line, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(line)
	st := conn.ConnectionState()
	fmt.Println(st.HandshakeComplete, tls.VersionName(st.Version), st.ServerName)
	// Output:
	// secure hello
	// true TLS 1.3 localhost
}

// Example_hostnameMismatch shows verification failing when the name is wrong.
func Example_hostnameMismatch() {
	cert, pool := selfSigned()
	ln, _ := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{cert}})
	defer ln.Close()
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		conn.(*tls.Conn).Handshake() // fails on both sides; we only care about the client
		conn.Close()
	}()

	_, err := tls.Dial("tcp", ln.Addr().String(), &tls.Config{
		RootCAs:    pool,
		ServerName: "other.example", // not in the certificate
	})
	var he x509.HostnameError
	fmt.Println(errors.As(err, &he))
	// Output: true
}

// Example_httpsServer serves HTTPS from an in-memory certificate and calls it
// with a client that trusts that certificate.
func Example_httpsServer() {
	cert, pool := selfSigned()
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello over https")
		}),
		TLSConfig:         &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12},
		ReadHeaderTimeout: 5 * time.Second,
	}
	go srv.ServeTLS(ln, "", "") // empty file names: use TLSConfig.Certificates
	defer srv.Close()

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12},
		},
	}
	resp, err := client.Get("https://" + ln.Addr().String() + "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	// Output: hello over https
}

// Example_config shows the readable names and safe defaults.
func Example_config() {
	cfg := &tls.Config{MinVersion: tls.VersionTLS12}
	fmt.Println(tls.VersionName(cfg.MinVersion), tls.VersionName(tls.VersionTLS13))
	fmt.Println(tls.CipherSuiteName(tls.TLS_AES_128_GCM_SHA256))
	fmt.Println(cfg.InsecureSkipVerify) // must stay false in real code
	// Output:
	// TLS 1.2 TLS 1.3
	// TLS_AES_128_GCM_SHA256
	// false
}
