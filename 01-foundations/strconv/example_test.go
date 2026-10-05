package strconvref_test

import (
	"errors"
	"fmt"
	"strconv"
)

// Example_atoiAndItoa covers the two most common conversions.
func Example_atoiAndItoa() {
	n, err := strconv.Atoi("123")
	fmt.Println(n, err)

	fmt.Println(strconv.Itoa(99) + "!")

	// Atoi is strict: no whitespace, no empty string, no decimals.
	for _, s := range []string{"abc", "", " 42", "3.5"} {
		_, err := strconv.Atoi(s)
		fmt.Println(err)
	}
	// Output:
	// 123 <nil>
	// 99!
	// strconv.Atoi: parsing "abc": invalid syntax
	// strconv.Atoi: parsing "": invalid syntax
	// strconv.Atoi: parsing " 42": invalid syntax
	// strconv.Atoi: parsing "3.5": invalid syntax
}

// Example_parseInt shows ParseInt(s, base, bitSize): the base, the number of
// bits the result must fit in, and how base 0 infers the base from a prefix.
func Example_parseInt() {
	v, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println(v)

	v, _ = strconv.ParseInt("-101", 2, 64)
	fmt.Println(v)

	v, _ = strconv.ParseInt("0x1f", 0, 64) // base 0: the 0x prefix selects base 16
	fmt.Println(v)

	v, _ = strconv.ParseInt("1_000", 0, 64) // underscores are allowed only with base 0
	fmt.Println(v)

	// A leading zero means octal when the base is 0.
	v, _ = strconv.ParseInt("010", 0, 64)
	fmt.Println(v)
	n, _ := strconv.Atoi("010") // Atoi always uses base 10
	fmt.Println(n)

	// bitSize 8 means "must fit in an int8". On overflow you get the nearest
	// limit plus an error.
	v, err := strconv.ParseInt("300", 10, 8)
	fmt.Println(v, err)

	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(u)
	// Output:
	// 255
	// -5
	// 31
	// 1000
	// 8
	// 10
	// 127 strconv.ParseInt: parsing "300": value out of range
	// 42
}

// Example_errorHandling shows how to inspect a parse failure.
func Example_errorHandling() {
	_, err := strconv.Atoi("abc")

	var ne *strconv.NumError
	if errors.As(err, &ne) {
		fmt.Println(ne.Func, ne.Num, ne.Err)
	}
	fmt.Println(errors.Is(err, strconv.ErrSyntax))

	_, err = strconv.ParseInt("999999999999999999999", 10, 64)
	fmt.Println(errors.Is(err, strconv.ErrRange))
	// Output:
	// Atoi abc invalid syntax
	// true
	// true
}

// Example_parseFloat shows float parsing, including the special values.
func Example_parseFloat() {
	for _, s := range []string{"3.14", "1e3", "inf", "NaN"} {
		f, err := strconv.ParseFloat(s, 64)
		fmt.Println(f, err)
	}
	_, err := strconv.ParseFloat("abc", 64)
	fmt.Println(err)
	// Output:
	// 3.14 <nil>
	// 1000 <nil>
	// +Inf <nil>
	// NaN <nil>
	// strconv.ParseFloat: parsing "abc": invalid syntax
}

// Example_formatFloat shows format ('f', 'e', 'g'), precision and bitSize.
// Precision -1 means "the fewest digits that round-trip exactly".
func Example_formatFloat() {
	fmt.Println(strconv.FormatFloat(3.14159, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(3.14159, 'e', 2, 64))
	fmt.Println(strconv.FormatFloat(1234567.0, 'e', 3, 64))
	fmt.Println(strconv.FormatFloat(1.0/3, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(2.5, 'f', -1, 64))

	// bitSize says which precision the shortest form is computed for.
	x := float64(float32(0.1))
	fmt.Println(strconv.FormatFloat(x, 'f', -1, 32))
	fmt.Println(strconv.FormatFloat(x, 'f', -1, 64))
	// Output:
	// 3.14
	// 3.14e+00
	// 1.235e+06
	// 0.3333333333333333
	// 2.5
	// 0.1
	// 0.10000000149011612
}

// Example_bool shows the exact set of strings ParseBool accepts.
func Example_bool() {
	for _, s := range []string{"1", "t", "TRUE", "false", "0"} {
		b, _ := strconv.ParseBool(s)
		fmt.Printf("%q -> %t\n", s, b)
	}
	_, err := strconv.ParseBool("yes")
	fmt.Println(err)
	fmt.Println(strconv.FormatBool(true))
	// Output:
	// "1" -> true
	// "t" -> true
	// "TRUE" -> true
	// "false" -> false
	// "0" -> false
	// strconv.ParseBool: parsing "yes": invalid syntax
	// true
}

// Example_formatInt converts integers to text in any base from 2 to 36.
func Example_formatInt() {
	fmt.Println(strconv.FormatInt(255, 2))
	fmt.Println(strconv.FormatInt(255, 16))
	fmt.Println(strconv.FormatInt(-255, 16))
	fmt.Println(strconv.FormatUint(255, 36))
	// Output:
	// 11111111
	// ff
	// -ff
	// 73
}

// Example_quote escapes strings the way Go source code would.
func Example_quote() {
	q := strconv.Quote("hi\n\"there\"")
	fmt.Println(q)

	s, err := strconv.Unquote(q)
	fmt.Printf("%q %v\n", s, err)

	fmt.Println(strconv.QuoteToASCII("héllo"))
	fmt.Println(strconv.QuoteRune('☺'))

	_, err = strconv.Unquote("abc") // input must already be quoted
	fmt.Println(err)
	// Output:
	// "hi\n\"there\""
	// "hi\n\"there\"" <nil>
	// "h\u00e9llo"
	// '☺'
	// invalid syntax
}

// Example_append builds output in a byte slice without intermediate strings.
func Example_append() {
	b := []byte("count=")
	b = strconv.AppendInt(b, 42, 10)
	b = append(b, ' ')
	b = strconv.AppendBool(b, true)
	b = append(b, ' ')
	b = strconv.AppendQuote(b, "go")
	fmt.Println(string(b))
	// Output:
	// count=42 true "go"
}

// Example_intToString shows the classic trap: converting an integer to a
// string with string(x) yields the character with that code, not the digits.
func Example_intToString() {
	i := 65
	fmt.Println(strconv.Itoa(i), string(rune(i)))
	// Output:
	// 65 A
}

// parsePort is a realistic use: parse, validate the range and add context.
func parsePort(s string) (int, error) {
	p, err := strconv.ParseUint(s, 10, 16) // 16 bits: 0 to 65535
	if err != nil {
		return 0, fmt.Errorf("invalid port %q: %w", s, err)
	}
	return int(p), nil
}

// Example_validation uses ParseUint's bitSize as free range checking and
// wraps the error with context.
func Example_validation() {
	for _, s := range []string{"8080", "70000", "-1"} {
		p, err := parsePort(s)
		fmt.Println(p, err)
	}
	// Output:
	// 8080 <nil>
	// 0 invalid port "70000": strconv.ParseUint: parsing "70000": value out of range
	// 0 invalid port "-1": strconv.ParseUint: parsing "-1": invalid syntax
}
