// Package solution holds the reference solutions for the embed exercises.
package solution

import (
	"embed"
	"path"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var tmplFS embed.FS

var tmpl = template.Must(template.ParseFS(tmplFS, "templates/*.tmpl"))

// Render executes the embedded template named name (for example "greeting")
// with the given person's name.
func Render(name, who string) (string, error) {
	var sb strings.Builder
	err := tmpl.ExecuteTemplate(&sb, name+".tmpl", map[string]string{"Name": who})
	return sb.String(), err
}

// Names returns the available template names without the .tmpl extension,
// in alphabetical order.
func Names() []string {
	entries, err := tmplFS.ReadDir("templates")
	if err != nil {
		return nil
	}
	var names []string
	for _, e := range entries {
		names = append(names, strings.TrimSuffix(e.Name(), path.Ext(e.Name())))
	}
	return names
}
