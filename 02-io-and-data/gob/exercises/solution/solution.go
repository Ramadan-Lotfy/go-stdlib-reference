// Package solution holds the reference solutions for the encoding/gob exercises.
package solution

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy returns a deep copy of src by encoding and decoding it with gob.
// Only exported fields are copied.
func DeepCopy[T any](src T) (T, error) {
	var dst T
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return dst, err
	}
	err := gob.NewDecoder(&buf).Decode(&dst)
	return dst, err
}
