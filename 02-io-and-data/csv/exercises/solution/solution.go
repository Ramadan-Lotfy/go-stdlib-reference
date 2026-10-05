// Package solution holds the reference solutions for the encoding/csv exercises.
package solution

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// User is one row of the input.
type User struct {
	Name string
	Age  int
}

// ParseUsers reads CSV with a header row containing "name" and "age" columns
// in any order and any letter case. A bad age is reported with its line number.
func ParseUsers(r io.Reader) ([]User, error) {
	cr := csv.NewReader(r)
	header, err := cr.Read()
	if err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}
	idx := make(map[string]int)
	for i, h := range header {
		idx[strings.ToLower(strings.TrimSpace(h))] = i
	}
	ni, ok1 := idx["name"]
	ai, ok2 := idx["age"]
	if !ok1 || !ok2 {
		return nil, errors.New("header must contain name and age")
	}

	var users []User
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			return users, nil
		}
		if err != nil {
			return nil, err
		}
		age, err := strconv.Atoi(strings.TrimSpace(rec[ai]))
		if err != nil {
			line, _ := cr.FieldPos(ai)
			return nil, fmt.Errorf("line %d: invalid age: %w", line, err)
		}
		users = append(users, User{Name: strings.TrimSpace(rec[ni]), Age: age})
	}
}
