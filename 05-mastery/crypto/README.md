# crypto/*

## Overview
The `crypto` packages provide hashing, message authentication, random numbers,
encryption and digital signatures. Go's implementations are carefully reviewed
and constant-time where it matters, but cryptography is easy to misuse.
Rule of thumb: pick a high-level, well-known construction (SHA-256, HMAC,
AES-GCM, Ed25519), never invent your own, and never encrypt with raw AES blocks.

| You want to... | Use |
|----------------|-----|
| Fingerprint data / check a download | `crypto/sha256` |
| Authenticate a message with a shared secret | `crypto/hmac` + SHA-256 |
| Generate keys, tokens, nonces | `crypto/rand` (never `math/rand`) |
| Encrypt and authenticate data | AES-GCM (`crypto/aes` + `crypto/cipher`) |
| Sign and verify with key pairs | `crypto/ed25519` |
| Compare secrets | `crypto/subtle.ConstantTimeCompare` |
| Hash passwords | `golang.org/x/crypto/bcrypt` or `argon2` (not in the stdlib, and **not** SHA-256) |
| Talk TLS | `crypto/tls` (see Stage 4) |

## Core API

| Function / Type | Purpose |
|---|---|
| `sha256.Sum256`, `sha256.New()`, `sha512`, `sha3` (1.24) | Cryptographic hashes; `New()` streams large data |
| `hmac.New(sha256.New, key)`, `hmac.Equal` | Keyed message authentication code |
| `rand.Read`, `rand.Int`, `rand.Reader` | Cryptographically secure random bytes and numbers |
| `aes.NewCipher`, `cipher.NewGCM` | Authenticated encryption: `Seal` and `Open` |
| `ed25519.GenerateKey`, `Sign`, `Verify` | Digital signatures |
| `subtle.ConstantTimeCompare` | Timing-safe comparison |
| `md5`, `sha1` | Broken for security; only for legacy formats and non-security checksums |
| `rsa`, `ecdsa`, `elliptic`, `x509` | Classic public-key crypto and certificates |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./05-mastery/crypto/
```

## Gotchas
- **Use `crypto/rand`, never `math/rand`,** for anything secret (tokens, keys, nonces). `math/rand` is predictable.
- **Never compare secrets or MACs with `==`.** It leaks timing. Use `hmac.Equal` or `subtle.ConstantTimeCompare`.
- **Do not store passwords with SHA-256 (or any fast hash).** Use bcrypt, scrypt or argon2id, which are deliberately slow and salted.
- **A GCM nonce must never repeat with the same key.** Generate it randomly (12 bytes) for each message and send it along with the ciphertext (it is not secret).
- **Encryption without authentication is dangerous.** Prefer AEAD modes (`GCM`, ChaCha20-Poly1305) that detect tampering; never use ECB or unauthenticated CBC.
- **Check every error,** especially from `Open` (tampered data) and `rand.Read`.
- **MD5 and SHA-1 are broken** for security (collisions). They remain fine for non-adversarial checksums and legacy protocols.
- **Keys are bytes, not passphrases.** Derive a key from a password with a KDF (argon2, scrypt, PBKDF2), do not use the password directly.
- **Do not implement your own protocol** (key exchange, session tokens, JWT parsing). Use well-reviewed libraries or `crypto/tls`.
- **Keep secrets out of logs, error messages and source code.**

## Java comparison

| Java | Go |
|---|---|
| `MessageDigest.getInstance("SHA-256")` | `sha256.New()` / `sha256.Sum256` |
| `Mac.getInstance("HmacSHA256")` | `hmac.New(sha256.New, key)` |
| `SecureRandom` | `crypto/rand` |
| `Cipher.getInstance("AES/GCM/NoPadding")` | `cipher.NewGCM(block)` |
| `Signature` (Ed25519) | `ed25519.Sign` / `Verify` |
| `MessageDigest.isEqual` | `subtle.ConstantTimeCompare` |
| Spring Security `BCryptPasswordEncoder` | `golang.org/x/crypto/bcrypt` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/crypto/*
- Source: https://cs.opensource.google/go/go/+/master:src/crypto/*/
