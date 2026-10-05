package cryptoref_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"strings"
)

// Example_sha256 hashes data in one call and by streaming.
func Example_sha256() {
	sum := sha256.Sum256([]byte("hello"))
	fmt.Printf("%x\n", sum)

	h := sha256.New() // streaming: feed any io.Writer source
	io.Copy(h, strings.NewReader("hel"))
	io.WriteString(h, "lo")
	fmt.Println(hex.EncodeToString(h.Sum(nil)) == hex.EncodeToString(sum[:]))
	fmt.Println(h.Size(), h.BlockSize())
	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
	// true
	// 32 64
}

// Example_legacyHashes shows MD5 and SHA-1, which are broken for security
// but still appear in old file formats and checksums.
func Example_legacyHashes() {
	fmt.Printf("%x\n", md5.Sum([]byte("abc")))
	fmt.Printf("%x\n", sha1.Sum([]byte("abc")))
	// Output:
	// 900150983cd24fb0d6963f7d28e17f72
	// a9993e364706816aba3e25717850c26c9cd0d89d
}

// Example_hmac authenticates a message with a shared key and verifies it.
func Example_hmac() {
	key := []byte("secret")
	sign := func(msg string) []byte {
		mac := hmac.New(sha256.New, key)
		mac.Write([]byte(msg))
		return mac.Sum(nil)
	}

	sig := sign("payload")
	fmt.Println(len(sig))
	fmt.Println(hmac.Equal(sig, sign("payload"))) // constant-time compare
	fmt.Println(hmac.Equal(sig, sign("payl0ad")))
	// Output:
	// 32
	// true
	// false
}

// Example_random generates secure random data.
func Example_random() {
	b := make([]byte, 16)
	n, err := rand.Read(b)
	fmt.Println(n, err)

	v, _ := rand.Int(rand.Reader, big.NewInt(100)) // uniform in [0, 100)
	fmt.Println(v.Sign() >= 0 && v.Cmp(big.NewInt(100)) < 0)
	// Output:
	// 16 <nil>
	// true
}

// Example_aesGCM encrypts and authenticates, then detects tampering.
func Example_aesGCM() {
	key := make([]byte, 32) // AES-256
	rand.Read(key)
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize()) // 12 bytes, unique per message
	rand.Read(nonce)
	sealed := gcm.Seal(nonce, nonce, []byte("attack at dawn"), nil) // nonce || ciphertext || tag

	ns := gcm.NonceSize()
	plain, err := gcm.Open(nil, sealed[:ns], sealed[ns:], nil)
	fmt.Println(string(plain), err)

	sealed[len(sealed)-1] ^= 1 // flip one bit
	_, err = gcm.Open(nil, sealed[:ns], sealed[ns:], nil)
	fmt.Println(err)
	// Output:
	// attack at dawn <nil>
	// cipher: message authentication failed
}

// Example_ed25519 signs with a private key and verifies with the public key.
func Example_ed25519() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	msg := []byte("release v1.2.3")
	sig := ed25519.Sign(priv, msg)

	fmt.Println(len(pub), len(sig))
	fmt.Println(ed25519.Verify(pub, msg, sig))
	fmt.Println(ed25519.Verify(pub, []byte("release v9.9.9"), sig))
	// Output:
	// 32 64
	// true
	// false
}

// Example_constantTime compares secrets without leaking where they differ.
func Example_constantTime() {
	fmt.Println(subtle.ConstantTimeCompare([]byte("token"), []byte("token")))
	fmt.Println(subtle.ConstantTimeCompare([]byte("token"), []byte("tokex")))
	fmt.Println(subtle.ConstantTimeCompare([]byte("token"), []byte("tok"))) // different length
	// Output:
	// 1
	// 0
	// 0
}
