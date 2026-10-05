package htmltemplateref_test

import (
	"fmt"
	htmltemplate "html/template"
	"os"
	texttemplate "text/template"
)

// Example_autoEscape contrasts html/template with text/template.
func Example_autoEscape() {
	const tpl = "<p>{{.}}</p>\n"
	input := "<script>alert('x')</script>"

	htmltemplate.Must(htmltemplate.New("h").Parse(tpl)).Execute(os.Stdout, input)
	texttemplate.Must(texttemplate.New("t").Parse(tpl)).Execute(os.Stdout, input)
	// Output:
	// <p>&lt;script&gt;alert(&#39;x&#39;)&lt;/script&gt;</p>
	// <p><script>alert('x')</script></p>
}

// Example_context shows that the same kind of value is escaped differently
// in a URL, an attribute and element text.
func Example_context() {
	t := htmltemplate.Must(htmltemplate.New("c").Parse(
		`<a href="/search?q={{.Q}}" title="{{.Label}}">{{.Label}}</a>` + "\n"))
	t.Execute(os.Stdout, map[string]string{"Q": "a b&c", "Label": `<Tom & "Jerry">`})
	// Output:
	// <a href="/search?q=a%20b%26c" title="&lt;Tom &amp; &#34;Jerry&#34;&gt;">&lt;Tom &amp; &#34;Jerry&#34;&gt;</a>
}

// Example_javascript embeds data into a script safely as a JavaScript value.
func Example_javascript() {
	t := htmltemplate.Must(htmltemplate.New("j").Parse(
		`<script>var user = {{.}};</script>` + "\n"))
	t.Execute(os.Stdout, map[string]any{"name": "Ada", "id": 7})
	// Output:
	// <script>var user = {"id":7,"name":"Ada"};</script>
}

// Example_trusted shows template.HTML, which bypasses escaping. Only ever use
// it for markup you wrote yourself, never for user input.
func Example_trusted() {
	t := htmltemplate.Must(htmltemplate.New("t").Parse("<p>{{.}}</p>\n"))
	t.Execute(os.Stdout, htmltemplate.HTML("<b>bold</b>"))
	// Output: <p><b>bold</b></p>
}

// Example_layout composes a page from a base layout and a page-specific block.
func Example_layout() {
	const base = `{{define "page"}}<h1>{{template "title" .}}</h1>{{block "body" .}}<p>default</p>{{end}}{{end}}`
	const home = `{{define "title"}}Home{{end}}{{define "body"}}<p>Hello {{.Name}}</p>{{end}}`

	t := htmltemplate.Must(htmltemplate.New("layout").Parse(base))
	htmltemplate.Must(t.Parse(home))

	err := t.ExecuteTemplate(os.Stdout, "page", map[string]string{"Name": "<Ada>"})
	fmt.Println()
	if err != nil {
		fmt.Println(err)
	}
	// Output: <h1>Home</h1><p>Hello &lt;Ada&gt;</p>
}
