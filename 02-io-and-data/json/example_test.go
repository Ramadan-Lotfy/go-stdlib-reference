package jsonref_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"-"`
	Tags     []string `json:"tags"`
	internal int
}

// Example_marshal shows tags: rename, omitempty, skip, unexported fields, nil slices.
func Example_marshal() {
	u := User{ID: 1, Name: "Ada", Password: "secret"}
	b, err := json.Marshal(u)
	fmt.Println(string(b), err)

	u.Tags = []string{} // empty, not nil
	b, _ = json.Marshal(u)
	fmt.Println(string(b))
	// Output:
	// {"id":1,"name":"Ada","tags":null} <nil>
	// {"id":1,"name":"Ada","tags":[]}
}

// Example_marshalIndent pretty-prints. Map keys are sorted.
func Example_marshalIndent() {
	b, _ := json.MarshalIndent(map[string]any{"b": []int{1, 2}, "a": "x"}, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//   "a": "x",
	//   "b": [
	//     1,
	//     2
	//   ]
	// }
}

// Example_unmarshal decodes into a struct (keys match case-insensitively,
// unknown keys are ignored) and into an untyped value.
func Example_unmarshal() {
	var u User
	err := json.Unmarshal([]byte(`{"ID": 7, "name": "Bob", "extra": true}`), &u)
	fmt.Printf("%+v %v\n", u, err)

	var v any
	json.Unmarshal([]byte(`{"a":1,"b":[true,null,"x"]}`), &v)
	m := v.(map[string]any)
	fmt.Printf("%T %v\n", m["a"], m["a"]) // numbers are float64
	fmt.Printf("%T %v\n", m["b"], m["b"])
	// Output:
	// {ID:7 Name:Bob Email: Password: Tags:[] internal:0} <nil>
	// float64 1
	// []interface {} [true <nil> x]
}

// Example_errors distinguishes syntax errors from type mismatches.
func Example_errors() {
	var u User
	err := json.Unmarshal([]byte(`{"id":}`), &u)
	var se *json.SyntaxError
	fmt.Println(errors.As(err, &se), err)

	err = json.Unmarshal([]byte(`{"id":"x"}`), &u)
	var te *json.UnmarshalTypeError
	fmt.Println(errors.As(err, &te), err)
	// Output:
	// true invalid character '}' looking for beginning of value
	// true json: cannot unmarshal string into Go struct field User.id of type int
}

// Example_decoder reads a stream of JSON values and can reject unknown fields.
func Example_decoder() {
	dec := json.NewDecoder(strings.NewReader(`{"n":1} {"n":2} {"n":3}`))
	var got []int
	for {
		var v struct {
			N int `json:"n"`
		}
		err := dec.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		got = append(got, v.N)
	}
	fmt.Println(got)

	strict := json.NewDecoder(strings.NewReader(`{"id":1,"extra":2}`))
	strict.DisallowUnknownFields()
	var u User
	fmt.Println(strict.Decode(&u))
	// Output:
	// [1 2 3]
	// json: unknown field "extra"
}

type Status int

const (
	Inactive Status = iota
	Active
)

func (s Status) String() string {
	if s == Active {
		return "active"
	}
	return "inactive"
}

func (s Status) MarshalJSON() ([]byte, error) { return json.Marshal(s.String()) }

func (s *Status) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	switch str {
	case "active":
		*s = Active
	case "inactive":
		*s = Inactive
	default:
		return fmt.Errorf("unknown status %q", str)
	}
	return nil
}

// Example_customMarshal gives a type its own JSON form.
func Example_customMarshal() {
	type account struct {
		S Status `json:"status"`
	}
	b, _ := json.Marshal(account{Active})
	fmt.Println(string(b))

	var a account
	err := json.Unmarshal([]byte(`{"status":"inactive"}`), &a)
	fmt.Println(a.S, err)

	err = json.Unmarshal([]byte(`{"status":"zzz"}`), &a)
	fmt.Println(err)
	// Output:
	// {"status":"active"}
	// inactive <nil>
	// unknown status "zzz"
}

// Example_rawMessage decodes the payload only after looking at the type field.
func Example_rawMessage() {
	type event struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}
	var e event
	json.Unmarshal([]byte(`{"type":"click","payload":{"x":1,"y":2}}`), &e)

	switch e.Type {
	case "click":
		var p struct{ X, Y int }
		json.Unmarshal(e.Payload, &p)
		fmt.Println("click at", p.X, p.Y)
	}
	// Output: click at 1 2
}

// Example_encoder writes to a stream. Encode adds a newline and escapes HTML by default.
func Example_encoder() {
	b, _ := json.Marshal("<b>&</b>")
	fmt.Println(string(b))

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.Encode("<b>&</b>")
	// Output:
	// "\u003cb\u003e\u0026\u003c/b\u003e"
	// "<b>&</b>"
}

// Example_numbers shows float64 precision loss and the UseNumber fix.
func Example_numbers() {
	data := `{"id":12345678901234567890}`

	var v map[string]any
	json.Unmarshal([]byte(data), &v)
	fmt.Println(v["id"])

	dec := json.NewDecoder(strings.NewReader(data))
	dec.UseNumber()
	var w map[string]any
	dec.Decode(&w)
	fmt.Println(w["id"])
	// Output:
	// 1.2345678901234567e+19
	// 12345678901234567890
}

// Example_time shows that time.Time uses RFC 3339.
func Example_time() {
	t := time.Date(2026, time.October, 3, 14, 30, 0, 0, time.UTC)
	b, _ := json.Marshal(map[string]time.Time{"at": t})
	fmt.Println(string(b))
	// Output: {"at":"2026-10-03T14:30:00Z"}
}
