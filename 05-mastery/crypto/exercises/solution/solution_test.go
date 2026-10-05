package solution

import "testing"

func TestSignKnownVector(t *testing.T) {
	// RFC 4231, test case 2.
	got := Sign([]byte("Jefe"), []byte("what do ya want for nothing?"))
	want := "5bdcc146bf60754e6a042426089575c75a003f089d2739839dec58b964ec3843"
	if got != want {
		t.Errorf("Sign() = %s\nwant   %s", got, want)
	}
}

func TestVerify(t *testing.T) {
	secret, msg := []byte("k"), []byte("hello")
	sig := Sign(secret, msg)

	if !Verify(secret, msg, sig) {
		t.Error("valid signature rejected")
	}
	if Verify(secret, []byte("hellO"), sig) {
		t.Error("changed message accepted")
	}
	if Verify([]byte("other"), msg, sig) {
		t.Error("wrong key accepted")
	}
	for _, bad := range []string{"", "zz", sig[:10], sig + "00"} {
		if Verify(secret, msg, bad) {
			t.Errorf("Verify accepted bad signature %q", bad)
		}
	}
}

func TestRandomToken(t *testing.T) {
	a, err := RandomToken(16)
	b, _ := RandomToken(16)
	if err != nil || len(a) != 32 || a == b {
		t.Errorf("RandomToken(16) = %q, %q, %v", a, b, err)
	}
	if tok, _ := RandomToken(0); tok != "" {
		t.Errorf("RandomToken(0) = %q, want empty", tok)
	}
}
