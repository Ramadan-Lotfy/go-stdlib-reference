// Package solution holds the reference solutions for the html/template exercises.
package solution

import (
	"html/template"
	"strings"
)

var commentsTemplate = template.Must(template.New("comments").Parse(
	`<ul>{{range .}}<li>{{.}}</li>{{end}}</ul>`))

// RenderComments renders user comments as an HTML list. Any markup in a
// comment must be escaped.
func RenderComments(comments []string) (string, error) {
	var sb strings.Builder
	if err := commentsTemplate.Execute(&sb, comments); err != nil {
		return "", err
	}
	return sb.String(), nil
}
