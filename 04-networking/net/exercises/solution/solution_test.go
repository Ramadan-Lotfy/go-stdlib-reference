package solution

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestIsPrivateIP(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"10.1.2.3", true},
		{"172.16.0.1", true},
		{"192.168.0.1", true},
		{"127.0.0.1", true},
		{"fe80::1", true},
		{"8.8.8.8", false},
		{"2001:4860:4860::8888", false},
	}
	for _, tt := range tests {
		got, err := IsPrivateIP(tt.in)
		if err != nil || got != tt.want {
			t.Errorf("IsPrivateIP(%q) = %v, %v; want %v", tt.in, got, err, tt.want)
		}
	}
	if _, err := IsPrivateIP("nope"); err == nil {
		t.Error("invalid IP should be an error")
	}
}

func TestServe(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- Serve(ln) }()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	conn.SetDeadline(time.Now().Add(2 * time.Second))
	fmt.Fprint(conn, "hello\nworld\n")

	r := bufio.NewReader(conn)
	for _, want := range []string{"HELLO\n", "WORLD\n"} {
		got, err := r.ReadString('\n')
		if err != nil || got != want {
			t.Fatalf("got %q, %v; want %q", got, err, want)
		}
	}
	conn.Close()

	ln.Close()
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Serve returned %v after Close, want nil", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Serve did not return after the listener was closed")
	}
}
