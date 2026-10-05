package regexpref_test

import (
	"fmt"
	"regexp"
	"strings"
)

// Example_match tests a whole string and a substring.
func Example_match() {
	phone := regexp.MustCompile(`^\d{3}-\d{4}$`)
	fmt.Println(phone.MatchString("555-1234"), phone.MatchString("5551234"))

	// Without anchors the pattern matches anywhere.
	fmt.Println(regexp.MustCompile(`\d+`).MatchString("abc123def"))
	fmt.Println(regexp.MustCompile(`(?i)hello`).MatchString("HeLLo"))
	// Output:
	// true false
	// true
	// true
}

// Example_find returns text and positions of matches.
func Example_find() {
	re := regexp.MustCompile(`\w+@\w+\.com`)
	s := "mail ada@example.com or bob@test.com"
	fmt.Println(re.FindString(s))
	fmt.Println(re.FindAllString(s, -1))
	fmt.Println(re.FindAllString(s, 1))
	fmt.Println(regexp.MustCompile(`o w`).FindStringIndex("hello world"))
	fmt.Println(re.FindString("nothing here") == "")
	// Output:
	// ada@example.com
	// [ada@example.com bob@test.com]
	// [ada@example.com]
	// [4 7]
	// true
}

// Example_submatch extracts capture groups, including named ones.
func Example_submatch() {
	re := regexp.MustCompile(`(?P<user>\w+)@(?P<host>[\w.]+)`)
	m := re.FindStringSubmatch("mail ada@example.com now")
	fmt.Println(m)
	for i, name := range re.SubexpNames() {
		if name != "" {
			fmt.Println(name, m[i])
		}
	}
	fmt.Println(re.SubexpIndex("host"))
	// Output:
	// [ada@example.com ada example.com]
	// user ada
	// host example.com
	// 2
}

// Example_replace rewrites text with group references or a function.
func Example_replace() {
	date := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	fmt.Println(date.ReplaceAllString("2026-10-03", "$3/$2/$1"))

	named := regexp.MustCompile(`(?P<y>\d+)-(?P<m>\d+)`)
	fmt.Println(named.ReplaceAllString("2026-10", "${m}/${y}"))

	words := regexp.MustCompile(`\w+`)
	fmt.Println(words.ReplaceAllStringFunc("go is fun", strings.ToUpper))

	fmt.Println(words.ReplaceAllLiteralString("a b", "$1")) // no expansion
	// Output:
	// 03/10/2026
	// 10/2026
	// GO IS FUN
	// $1 $1
}

// Example_dollarGotcha shows why ${1} is safer than $1.
func Example_dollarGotcha() {
	re := regexp.MustCompile(`(a)`)
	fmt.Printf("%q\n", re.ReplaceAllString("a", "${1}x"))
	fmt.Printf("%q\n", re.ReplaceAllString("a", "$1x")) // group named "1x" does not exist
	// Output:
	// "ax"
	// ""
}

// Example_split splits on a pattern.
func Example_split() {
	re := regexp.MustCompile(`\s*,\s*`)
	fmt.Printf("%q\n", re.Split("a , b,c ,  d", -1))
	// Output: ["a" "b" "c" "d"]
}

// Example_compileErrors shows Compile vs MustCompile and QuoteMeta.
func Example_compileErrors() {
	_, err := regexp.Compile(`a(b`)
	fmt.Println(err)

	// Escape user text before embedding it in a pattern.
	literal := regexp.QuoteMeta("1+1=2")
	fmt.Println(literal)
	fmt.Println(regexp.MustCompile(literal).MatchString("so 1+1=2 it is"))
	// Output:
	// error parsing regexp: missing closing ): `a(b`
	// 1\+1=2
	// true
}
