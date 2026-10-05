# crypto/* exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./05-mastery/crypto/exercises/...
```

## Exercise: signed messages and tokens

Write three functions:

```go
func Sign(secret, msg []byte) string               // hex HMAC-SHA256
func Verify(secret, msg []byte, sigHex string) bool // constant-time check
func RandomToken(n int) (string, error)             // n random bytes as hex
```

Rules:

- `Sign` must match the RFC 4231 test vector in the tests (key `Jefe`, message
  `what do ya want for nothing?`)
- `Verify` returns `false` (it does not panic or error) for invalid hex,
  signatures of the wrong length, a changed message or a wrong key
- Compare with `hmac.Equal`, never with `==`
- `RandomToken` uses `crypto/rand`, so two calls differ, and returns `2n` hex
  characters

Hints: `hmac.New(sha256.New, key)`, `hex.EncodeToString`, `hex.DecodeString`.

## Bonus
Write `Seal(key, plaintext []byte) ([]byte, error)` and `Open(key, sealed []byte) ([]byte, error)`
with AES-GCM and a random nonce prepended to the output. Test that a flipped bit
makes `Open` fail.
