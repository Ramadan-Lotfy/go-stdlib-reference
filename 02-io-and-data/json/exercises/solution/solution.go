// Package solution holds the reference solutions for the encoding/json exercises.
package solution

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Config is the application configuration.
type Config struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Debug bool   `json:"debug"`
}

// ParseConfig decodes strict JSON: unknown fields and trailing data are
// errors, missing fields get defaults (localhost, 8080), and the port must be
// between 1 and 65535.
func ParseConfig(data []byte) (Config, error) {
	cfg := Config{Host: "localhost", Port: 8080}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	var extra json.RawMessage
	if err := dec.Decode(&extra); err != io.EOF {
		return Config{}, errors.New("parse config: trailing data")
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		return Config{}, fmt.Errorf("parse config: port %d out of range", cfg.Port)
	}
	return cfg, nil
}

// PrettyJSON re-indents JSON with two spaces.
func PrettyJSON(data []byte) (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		return "", err
	}
	return buf.String(), nil
}
