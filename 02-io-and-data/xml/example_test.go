package xmlref_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Authors []string `xml:"authors>author"`
}

// Example_marshal writes attributes, elements and a nested path.
func Example_marshal() {
	b, err := xml.MarshalIndent(Book{ID: "b1", Title: "Go", Authors: []string{"A", "B"}}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// <book id="b1">
	//   <title>Go</title>
	//   <authors>
	//     <author>A</author>
	//     <author>B</author>
	//   </authors>
	// </book>
}

// Example_unmarshal parses the same shape back.
func Example_unmarshal() {
	data := `<book id="b2"><title>Concurrency</title><authors><author>X</author></authors></book>`
	var b Book
	err := xml.Unmarshal([]byte(data), &b)
	fmt.Println(b.ID, b.Title, b.Authors, err)
	// Output: b2 Concurrency [X] <nil>
}

type Price struct {
	Currency string `xml:"currency,attr"`
	Amount   string `xml:",chardata"`
}

// Example_chardata mixes an attribute with text content. The element name
// defaults to the type name.
func Example_chardata() {
	b, _ := xml.Marshal(Price{"USD", "9.99"})
	fmt.Println(xml.Header + string(b))
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <Price currency="USD">9.99</Price>
}

// Example_tokens walks a document without building a tree.
func Example_tokens() {
	d := xml.NewDecoder(strings.NewReader(`<a><b>x</b><c/></a>`))
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		switch t := tok.(type) {
		case xml.StartElement:
			fmt.Println("start", t.Name.Local)
		case xml.EndElement:
			fmt.Println("end", t.Name.Local)
		case xml.CharData:
			fmt.Printf("text %q\n", string(t))
		}
	}
	// Output:
	// start a
	// start b
	// text "x"
	// end b
	// start c
	// end c
	// end a
}

// Example_escape escapes text for safe inclusion in XML.
func Example_escape() {
	xml.EscapeText(os.Stdout, []byte(`a<b & "c"`))
	fmt.Println()
	// Output: a&lt;b &amp; &#34;c&#34;
}

// Example_errors shows a malformed document.
func Example_errors() {
	var b Book
	err := xml.Unmarshal([]byte(`<book><title>oops</book>`), &b)
	fmt.Println(err)
	// Output: XML syntax error on line 1: element <title> closed by </book>
}
