// Package solution holds the reference solutions for the net/url exercises.
package solution

import (
	"errors"
	"net/url"
)

// AddQuery returns rawURL with the given parameters added or replaced. Existing
// parameters are kept. The URL must have a scheme and a host.
func AddQuery(rawURL string, params map[string]string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", errors.New("URL needs a scheme and a host")
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// Domain returns the host name of rawURL without port or credentials.
func Domain(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Hostname() == "" {
		return "", errors.New("no host in URL")
	}
	return u.Hostname(), nil
}
