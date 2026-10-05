// Package solution holds the reference solutions for the strconv exercises.
package solution

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseKV parses "a=1,b=2" into map[a:1 b:2]. Spaces around keys and values
// are ignored. An empty input gives an empty map.
func ParseKV(s string) (map[string]int, error) {
	result := make(map[string]int)
	if strings.TrimSpace(s) == "" {
		return result, nil
	}
	for _, pair := range strings.Split(s, ",") {
		key, val, ok := strings.Cut(pair, "=")
		if !ok {
			return nil, fmt.Errorf("invalid pair %q", strings.TrimSpace(pair))
		}
		key = strings.TrimSpace(key)
		n, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return nil, fmt.Errorf("invalid value for %q: %w", key, err)
		}
		result[key] = n
	}
	return result, nil
}

// HumanBytes formats a byte count using 1024-based units, with at most one
// decimal and no trailing ".0". n must be >= 0.
func HumanBytes(n int64) string {
	if n < 1024 {
		return strconv.FormatInt(n, 10) + " B"
	}
	units := [...]string{"KB", "MB", "GB", "TB"}
	v := float64(n) / 1024
	i := 0
	for {
		s := strconv.FormatFloat(v, 'f', 1, 64)
		// Stay in this unit unless the value is >= 1024, or it would print
		// as "1024.0" after rounding. The last unit has nowhere to promote to.
		if i == len(units)-1 || (v < 1024 && s != "1024.0") {
			return strings.TrimSuffix(s, ".0") + " " + units[i]
		}
		v /= 1024
		i++
	}
}
