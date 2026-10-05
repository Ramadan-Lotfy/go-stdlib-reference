// Package solution holds the reference solutions for the encoding/xml exercises.
package solution

import "encoding/xml"

type feed struct {
	Titles []string `xml:"entry>title"`
}

// FeedTitles returns the title of every <entry> in a document shaped like
//
//	<feed><entry><title>One</title></entry>...</feed>
func FeedTitles(data []byte) ([]string, error) {
	var f feed
	if err := xml.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	return f.Titles, nil
}
