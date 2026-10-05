package stringsref_test

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Example_searching covers the Contains, Index, Count and prefix/suffix family.
func Example_searching() {
	s := "gopher gopher go"
	fmt.Println(strings.Contains(s, "pher"))
	fmt.Println(strings.ContainsAny(s, "xyz"))
	fmt.Println(strings.ContainsRune(s, 'h'))
	fmt.Println(strings.HasPrefix(s, "go"), strings.HasSuffix(s, "go"))
	fmt.Println(strings.Index(s, "go"), strings.LastIndex(s, "go"))
	fmt.Println(strings.Index(s, "rust")) // -1 means "not found"
	fmt.Println(strings.Count(s, "go"))
	fmt.Println(strings.EqualFold("Go", "GO")) // case-insensitive compare
	// Output:
	// true
	// false
	// true
	// true true
	// 0 14
	// -1
	// 3
	// true
}

// Example_splitting shows Split, SplitN, Fields and the empty-string edge case.
// %q makes the element boundaries visible.
func Example_splitting() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 2))
	fmt.Printf("%q\n", strings.SplitAfter("a,b", ","))
	fmt.Printf("%q\n", strings.Split("", ",")) // one empty string, not zero elements
	fmt.Printf("%q\n", strings.Split("a b  c", " "))
	fmt.Printf("%q\n", strings.Fields("  a   b \t c\n")) // splits on runs of whitespace
	fmt.Printf("%q\n", strings.FieldsFunc("a1b22c", unicode.IsDigit))
	// Output:
	// ["a" "b" "c"]
	// ["a" "b,c,d"]
	// ["a," "b"]
	// [""]
	// ["a" "b" "" "c"]
	// ["a" "b" "c"]
	// ["a" "b" "c"]
}

// Example_joinAndRepeat builds strings from pieces.
func Example_joinAndRepeat() {
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println(strings.Repeat("ab", 3))
	fmt.Println(strings.Repeat("=", 10))
	// Output:
	// a-b-c
	// ababab
	// ==========
}

// Example_trimming contrasts cutset functions (Trim, TrimLeft) with the
// exact-match functions (TrimPrefix, TrimSuffix).
func Example_trimming() {
	fmt.Printf("%q\n", strings.TrimSpace("  hi \n"))
	fmt.Println(strings.Trim("xxhixx", "x"))
	fmt.Println(strings.TrimLeft("baab-x", "ab"))   // strips any of 'a' and 'b'
	fmt.Println(strings.TrimPrefix("baab-x", "ab")) // prefix "ab" is not there
	fmt.Println(strings.TrimSuffix("file.go", ".go"))
	fmt.Println(strings.TrimFunc("123abc456", unicode.IsDigit))
	// Output:
	// "hi"
	// hi
	// -x
	// baab-x
	// file
	// abc
}

// Example_cut splits around the first separator. It is usually cleaner than
// Index plus slicing.
func Example_cut() {
	before, after, found := strings.Cut("key=value=x", "=")
	fmt.Println(before, after, found)

	_, _, found = strings.Cut("novalue", "=")
	fmt.Println(found)

	token, ok := strings.CutPrefix("Bearer abc123", "Bearer ")
	fmt.Println(token, ok)
	// Output:
	// key value=x true
	// false
	// abc123 true
}

// Example_caseAndReplace covers case conversion, Replace and Map.
func Example_caseAndReplace() {
	fmt.Println(strings.ToUpper("go"), strings.ToLower("GO"))
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // at most 2 replacements
	fmt.Println(strings.ReplaceAll("oink oink", "oink", "moo"))

	upperA := func(r rune) rune {
		if r == 'a' {
			return 'A'
		}
		return r
	}
	fmt.Println(strings.Map(upperA, "banana"))

	dropSpaces := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // a negative value drops the rune
		}
		return r
	}
	fmt.Println(strings.Map(dropSpaces, "a b c"))
	// Output:
	// GO go
	// oinky oinky oink
	// moo moo
	// bAnAnA
	// abc
}

// Example_builder shows efficient string construction. Strings are immutable,
// so building one with += in a loop copies the data every time.
func Example_builder() {
	var sb strings.Builder
	sb.Grow(32) // optional: preallocate when you know the rough size
	sb.WriteString("go")
	sb.WriteByte('-')
	sb.WriteRune('✓')
	fmt.Fprintf(&sb, " v%d", 1) // Builder is an io.Writer
	fmt.Println(sb.String(), sb.Len())
	// Output:
	// go-✓ v1 9
}

// Example_replacer performs several replacements in a single pass.
func Example_replacer() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("<b>hi</b>"))
	// Output:
	// &lt;b&gt;hi&lt;/b&gt;
}

// Example_reader turns a string into an io.Reader so it can feed any API
// that reads streams.
func Example_reader() {
	data, _ := io.ReadAll(strings.NewReader("abc"))
	fmt.Println(string(data))

	sc := bufio.NewScanner(strings.NewReader("line1\nline2\nline3"))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	// Output:
	// abc
	// line1
	// line2
	// line3
}

// Example_unicode shows that strings are bytes: len counts bytes, while
// ranging over a string yields runes with their byte offsets.
func Example_unicode() {
	s := "héy"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	for i, r := range s {
		fmt.Printf("%d %c\n", i, r)
	}
	fmt.Println(strings.ToUpper("héllo"))
	fmt.Println(strings.IndexRune(s, 'y')) // byte offset, not rune offset
	// Output:
	// 4 3
	// 0 h
	// 1 é
	// 3 y
	// HÉLLO
	// 3
}
