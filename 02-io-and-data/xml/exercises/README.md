# encoding/xml exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./02-io-and-data/xml/exercises/...
```

## Exercise: FeedTitles

Write `FeedTitles(data []byte) ([]string, error)` that returns the title of
every entry in a document shaped like:

```xml
<feed>
  <entry><title>One</title></entry>
  <entry><title>Two &amp; Three</title></entry>
</feed>
```

Result: `["One", "Two & Three"]`. An empty feed returns an empty slice; malformed
XML returns an error.

Hints: a struct with a field tagged `xml:"entry>title"`, then `xml.Unmarshal`.

## Bonus
Rewrite it with `xml.Decoder.Token` so it never holds the whole document in
memory, and stop early after the first N titles.
