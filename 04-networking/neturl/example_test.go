package neturlref_test

import (
	"errors"
	"fmt"
	"net/url"
)

// Example_parse takes a URL apart.
func Example_parse() {
	u, err := url.Parse("https://user:pass@example.com:8443/a/b?q=go&q=lang&x=1#frag")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Scheme, u.Host, u.Hostname(), u.Port())
	fmt.Println(u.Path, u.RawQuery, u.Fragment)
	fmt.Println(u.User.Username())
	fmt.Println(u.Redacted())
	// Output:
	// https example.com:8443 example.com 8443
	// /a/b q=go&q=lang&x=1 frag
	// user
	// https://user:xxxxx@example.com:8443/a/b?q=go&q=lang&x=1#frag
}

// Example_query reads query parameters, including repeated ones.
func Example_query() {
	u, _ := url.Parse("https://example.com/s?q=go&q=lang&x=1")
	v := u.Query()
	fmt.Println(v.Get("q"), v["q"], v.Get("missing") == "", v.Has("x"))
	fmt.Println(v.Encode()) // keys are sorted
	// Output:
	// go [go lang] true true
	// q=go&q=lang&x=1
}

// Example_build assembles a URL from parts, escaping values correctly.
func Example_build() {
	u := url.URL{Scheme: "https", Host: "example.com", Path: "/search"}
	q := url.Values{}
	q.Set("q", "a b&c")
	q.Add("tag", "x")
	q.Add("tag", "y")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	// Output: https://example.com/search?q=a+b%26c&tag=x&tag=y
}

// Example_escape contrasts query and path escaping.
func Example_escape() {
	fmt.Println(url.QueryEscape("a b/c?"))
	fmt.Println(url.PathEscape("a b/c"))

	s, err := url.QueryUnescape("a%20b+c")
	fmt.Println(s, err)

	_, err = url.QueryUnescape("%zz")
	fmt.Println(err)
	// Output:
	// a+b%2Fc%3F
	// a%20b%2Fc
	// a b c <nil>
	// invalid URL escape "%zz"
}

// Example_resolve combines a base URL and a relative reference.
func Example_resolve() {
	base, _ := url.Parse("https://example.com/a/b/")
	ref, _ := url.Parse("../c?x=1")
	fmt.Println(base.ResolveReference(ref))

	joined, _ := url.JoinPath("https://example.com/api", "v1", "users")
	fmt.Println(joined)
	// Output:
	// https://example.com/a/c?x=1
	// https://example.com/api/v1/users
}

// Example_lenientParse shows that Parse is permissive and what to check instead.
func Example_lenientParse() {
	u, err := url.Parse("example.com/path")
	fmt.Printf("%q %q %q %v\n", u.Scheme, u.Host, u.Path, err)

	_, err = url.ParseRequestURI("example.com/path")
	fmt.Println(err != nil)

	_, err = url.Parse("http://[::1")
	var ue *url.Error
	fmt.Println(errors.As(err, &ue), ue.Op)
	// Output:
	// "" "" "example.com/path" <nil>
	// true
	// true parse
}
