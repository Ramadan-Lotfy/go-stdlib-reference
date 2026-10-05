// Package solution holds the reference solutions for the net exercises.
package solution

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

// Serve accepts connections on ln and, for each one, replies to every line
// with its upper-case version. It returns nil when ln is closed.
func Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Fprintln(conn, strings.ToUpper(sc.Text()))
	}
}

// IsPrivateIP reports whether s is a private, loopback or link-local address.
func IsPrivateIP(s string) (bool, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return false, fmt.Errorf("invalid IP %q", s)
	}
	return ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast(), nil
}
