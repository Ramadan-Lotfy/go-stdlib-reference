// Package solution holds the reference solutions for the crypto/* exercises.
package solution

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Sign returns the hex-encoded HMAC-SHA256 of msg under secret.
func Sign(secret, msg []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write(msg)
	return hex.EncodeToString(mac.Sum(nil))
}

// Verify reports whether sigHex is the valid signature of msg, comparing in
// constant time. Invalid hex is simply "not valid".
func Verify(secret, msg []byte, sigHex string) bool {
	got, err := hex.DecodeString(sigHex)
	if err != nil {
		return false
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write(msg)
	return hmac.Equal(got, mac.Sum(nil))
}

// RandomToken returns n random bytes as a hex string (2n characters).
func RandomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
