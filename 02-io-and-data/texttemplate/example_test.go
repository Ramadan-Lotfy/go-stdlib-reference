package texttemplateref_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

// Example_basic fills fields and calls the built-in len.
func Example_basic() {
	t := template.Must(template.New("greet").Parse("Hello, {{.Name}}! You have {{len .Items}} items.\n"))
	err := t.Execute(os.Stdout, struct {
		Name  string
		Items []string
	}{"Ada", []string{"a", "b"}})
	if err != nil {
		fmt.Println(err)
	}
	// Output: Hello, Ada! You have 2 items.
}

// Example_range loops with an index, a separator and an else branch for empty input.
func Example_range() {
	t := template.Must(template.New("r").Parse(
		`{{range $i, $e := .}}{{if $i}}, {{end}}{{$i}}={{$e}}{{else}}empty{{end}}` + "\n"))
	t.Execute(os.Stdout, []string{"a", "b", "c"})
	t.Execute(os.Stdout, []string{})
	// Output:
	// 0=a, 1=b, 2=c
	// empty
}

// Example_trim uses {{- to remove whitespace so a list renders cleanly.
func Example_trim() {
	t := template.Must(template.New("t").Parse("{{- range . }}\n- {{ . }}\n{{- end }}\n"))
	t.Execute(os.Stdout, []string{"x", "y"})
	// Output:
	// - x
	// - y
}

// Example_funcs registers functions and chains them with a pipeline.
func Example_funcs() {
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
		"add":   func(a, b int) int { return a + b },
	}
	t := template.Must(template.New("f").Funcs(funcs).Parse(
		`{{.Name | upper}} {{add 2 3}} {{printf "%05.1f" .Score}}` + "\n"))
	t.Execute(os.Stdout, map[string]any{"Name": "ada", "Score": 3.14159})
	// Output: ADA 5 003.1
}

// Example_conditions combines comparison and logic functions.
func Example_conditions() {
	t := template.Must(template.New("c").Parse(
		`{{if and (gt .Age 17) (eq .Role "admin")}}allowed{{else}}denied{{end}}` + "\n"))
	t.Execute(os.Stdout, struct {
		Age  int
		Role string
	}{20, "admin"})
	t.Execute(os.Stdout, struct {
		Age  int
		Role string
	}{20, "guest"})
	// Output:
	// allowed
	// denied
}

// Example_define reuses a named template inside a loop.
func Example_define() {
	const text = `{{define "item"}}<{{.}}>{{end}}{{range .}}{{template "item" .}}{{end}}` + "\n"
	t := template.Must(template.New("list").Parse(text))
	t.Execute(os.Stdout, []string{"a", "b"})
	// Output: <a><b>
}

// Example_missingKey makes a missing map key an error instead of "<no value>".
func Example_missingKey() {
	lenient := template.Must(template.New("m").Parse("[{{.name}}]\n"))
	lenient.Execute(os.Stdout, map[string]any{})

	strict := template.Must(template.New("m").Option("missingkey=error").Parse("[{{.name}}]"))
	err := strict.Execute(io.Discard, map[string]any{})
	fmt.Println(err)
	// Output:
	// [<no value>]
	// template: m:1:3: executing "m" at <.name>: map has no entry for key "name"
}
