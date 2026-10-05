// Package solution holds the reference solutions for the net/http/httptest exercises.
package solution

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetJSON performs a GET request and decodes a JSON response into v.
// It returns an error for transport failures, any status other than 200, and
// invalid JSON. Responses are limited to 1 MiB.
func GetJSON(ctx context.Context, client *http.Client, url string, v any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: unexpected status %d", url, resp.StatusCode)
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(v); err != nil {
		return fmt.Errorf("decode %s: %w", url, err)
	}
	return nil
}
