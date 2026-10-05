package csvref_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Example_readAll parses quoted fields that contain commas.
func Example_readAll() {
	in := "name,age\nAda,36\n\"Smith, J\",41\n"
	records, err := csv.NewReader(strings.NewReader(in)).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, rec := range records {
		fmt.Printf("%q\n", rec)
	}
	// Output:
	// ["name" "age"]
	// ["Ada" "36"]
	// ["Smith, J" "41"]
}

// Example_readLoop streams records and shows the field-count check.
func Example_readLoop() {
	r := csv.NewReader(strings.NewReader("a,b\nc\n"))
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		fmt.Println(rec)
	}
	// Output:
	// [a b]
	// error: record on line 2: wrong number of fields
}

// Example_dialect reads semicolon-separated data with comments.
func Example_dialect() {
	r := csv.NewReader(strings.NewReader("# a comment\nx;y\n1;2\n"))
	r.Comma = ';'
	r.Comment = '#'
	records, _ := r.ReadAll()
	fmt.Println(records)
	// Output: [[x y] [1 2]]
}

// Example_ragged allows rows with different numbers of fields.
func Example_ragged() {
	r := csv.NewReader(strings.NewReader("a,b,c\nd\n"))
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	fmt.Println(records, err)
	// Output: [[a b c] [d]] <nil>
}

// Example_writer quotes fields automatically.
func Example_writer() {
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"name", "note"})
	w.Write([]string{"Ada", `likes "quotes", commas`})
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
	}
	// Output:
	// name,note
	// Ada,"likes ""quotes"", commas"
}

// Example_fieldPos reports where a field is, which helps with error messages.
func Example_fieldPos() {
	r := csv.NewReader(strings.NewReader("a,b\nc,d\n"))
	r.Read()
	r.Read()
	line, col := r.FieldPos(1)
	fmt.Println(line, col)
	// Output: 2 3
}
