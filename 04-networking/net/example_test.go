package netref_test

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Example_parseIP classifies addresses.
func Example_parseIP() {
	ip := net.ParseIP("192.168.1.10")
	fmt.Println(ip, ip.To4() != nil, ip.IsPrivate(), ip.IsLoopback())

	v6 := net.ParseIP("::1")
	fmt.Println(v6, v6.To4() == nil, v6.IsLoopback())

	fmt.Println(net.ParseIP("not-an-ip") == nil)
	// Output:
	// 192.168.1.10 true true false
	// ::1 true true
	// true
}

// Example_cidr tests subnet membership.
func Example_cidr() {
	ip, ipnet, _ := net.ParseCIDR("10.0.0.5/24")
	fmt.Println(ip, ipnet)

	ones, bits := ipnet.Mask.Size()
	fmt.Println(ones, bits)

	fmt.Println(ipnet.Contains(net.ParseIP("10.0.0.77")))
	fmt.Println(ipnet.Contains(net.ParseIP("10.0.1.1")))
	// Output:
	// 10.0.0.5 10.0.0.0/24
	// 24 32
	// true
	// false
}

// Example_hostPort splits and joins addresses safely, including IPv6.
func Example_hostPort() {
	host, port, err := net.SplitHostPort("example.com:8080")
	fmt.Println(host, port, err)

	fmt.Println(net.JoinHostPort("::1", "80"))

	_, _, err = net.SplitHostPort("noport")
	fmt.Println(err)
	// Output:
	// example.com 8080 <nil>
	// [::1]:80
	// address noport: missing port in address
}

// Example_tcpEcho runs a tiny TCP echo server and talks to it.
func Example_tcpEcho() {
	ln, err := net.Listen("tcp", "127.0.0.1:0") // port 0: pick a free port
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
		io.Copy(conn, conn) // echo until the client closes
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Fprint(conn, "hello\n")
	line, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("%q\n", line)
	// Output: "hello\n"
}

// Example_pipe uses net.Pipe, an in-memory Conn pair that is ideal for tests.
func Example_pipe() {
	client, server := net.Pipe()
	go func() {
		fmt.Fprint(server, "ping")
		server.Close()
	}()
	data, _ := io.ReadAll(client)
	fmt.Println(string(data))
	// Output: ping
}

// Example_deadline shows a read deadline firing on a silent peer.
func Example_deadline() {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	client.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	_, err := client.Read(make([]byte, 1))

	var ne net.Error
	fmt.Println(errors.As(err, &ne) && ne.Timeout())
	fmt.Println(errors.Is(err, os.ErrDeadlineExceeded))
	// Output:
	// true
	// true
}

// Example_udp sends one datagram.
func Example_udp() {
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pc.Close()

	conn, _ := net.Dial("udp", pc.LocalAddr().String())
	defer conn.Close()
	conn.Write([]byte("hi"))

	buf := make([]byte, 16)
	pc.SetReadDeadline(time.Now().Add(time.Second))
	n, _, err := pc.ReadFrom(buf)
	fmt.Println(string(buf[:n]), err)
	// Output: hi <nil>
}

// Example_dialError inspects a failed connection attempt.
func Example_dialError() {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	addr := ln.Addr().String()
	ln.Close() // nobody listens here any more

	_, err := net.DialTimeout("tcp", addr, time.Second)
	var oe *net.OpError
	if errors.As(err, &oe) {
		fmt.Println(oe.Op, oe.Net)
	}
	// Output: dial tcp
}
