package osexecref_test

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const helperEnv = "OSEXEC_HELPER"

// TestMain makes this test binary double as the child program. When
// OSEXEC_HELPER=1 the binary runs the helper command instead of the tests.
func TestMain(m *testing.M) {
	if os.Getenv(helperEnv) == "1" {
		os.Exit(helper(os.Args[1:]))
	}
	os.Exit(m.Run())
}

func helper(args []string) int {
	switch args[0] {
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "upper":
		data, _ := io.ReadAll(os.Stdin)
		fmt.Print(strings.ToUpper(string(data)))
	case "fail":
		fmt.Fprintln(os.Stderr, "something broke")
		return 3
	case "both":
		fmt.Fprint(os.Stdout, "out;")
		fmt.Fprint(os.Stderr, "err;")
	case "env":
		fmt.Println(os.Getenv("GREETING"))
	case "lines":
		for i := 1; i <= 3; i++ {
			fmt.Printf("line %d\n", i)
		}
	case "sleep":
		time.Sleep(time.Minute)
	}
	return 0
}

// child builds a command that runs this binary in helper mode.
func child(args ...string) *exec.Cmd {
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), helperEnv+"=1")
	return cmd
}

// Example_output captures a command's standard output.
func Example_output() {
	out, err := child("echo", "hello", "world").Output()
	fmt.Printf("%q %v\n", out, err)
	// Output: "hello world\n" <nil>
}

// Example_noShell shows that arguments are passed verbatim: there is no shell
// to interpret ';' or '$HOME'.
func Example_noShell() {
	out, _ := child("echo", "; rm -rf /", "$HOME").Output()
	fmt.Printf("%q\n", out)
	// Output: "; rm -rf / $HOME\n"
}

// Example_stdin feeds input to the command.
func Example_stdin() {
	cmd := child("upper")
	cmd.Stdin = strings.NewReader("hello")
	out, err := cmd.Output()
	fmt.Printf("%q %v\n", out, err)
	// Output: "HELLO" <nil>
}

// Example_exitCode inspects a failing command.
func Example_exitCode() {
	cmd := child("fail")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		fmt.Println(ee.ExitCode(), strings.TrimSpace(stderr.String()))
	}
	// Output: 3 something broke
}

// Example_combined captures stdout and stderr together.
func Example_combined() {
	out, _ := child("both").CombinedOutput()
	s := string(out)
	fmt.Println(strings.Contains(s, "out;"), strings.Contains(s, "err;"))
	// Output: true true
}

// Example_env passes environment variables to the child.
func Example_env() {
	cmd := child("env")
	cmd.Env = append(cmd.Env, "GREETING=hi")
	out, _ := cmd.Output()
	fmt.Printf("%q\n", out)
	// Output: "hi\n"
}

// Example_streaming reads output while the command is still running.
func Example_streaming() {
	cmd := child("lines")
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	sc := bufio.NewScanner(stdout)
	for sc.Scan() {
		fmt.Println("got:", sc.Text())
	}
	fmt.Println(cmd.Wait()) // always Wait after Start
	// Output:
	// got: line 1
	// got: line 2
	// got: line 3
	// <nil>
}

// Example_timeout kills a command that runs too long.
func Example_timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, os.Args[0], "sleep")
	cmd.Env = append(os.Environ(), helperEnv+"=1")
	cmd.WaitDelay = time.Second // do not wait forever for a stuck process

	err := cmd.Run()
	fmt.Println(err != nil, errors.Is(ctx.Err(), context.DeadlineExceeded))
	// Output: true true
}

// Example_notFound shows the error for a program that does not exist.
func Example_notFound() {
	_, err := exec.LookPath("definitely-not-a-real-program")
	fmt.Println(errors.Is(err, exec.ErrNotFound))

	err = exec.Command("definitely-not-a-real-program").Run()
	fmt.Println(errors.Is(err, exec.ErrNotFound))
	// Output:
	// true
	// true
}
