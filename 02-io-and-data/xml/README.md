# encoding/xml

## Overview
`encoding/xml` marshals structs to XML and parses XML into structs, and offers
a token-level `Decoder` for large documents. Struct tags control element names,
attributes, character data and nesting.

## Core API

| Function / Type | Purpose |
|---|---|
| `xml.Marshal`, `xml.MarshalIndent` | Struct to XML |
| `xml.Unmarshal` | XML into a struct |
| `xml:"name"`, `xml:"id,attr"`, `xml:",chardata"` | Element, attribute, text content |
| `xml:"a>b"` | Nested path without an extra struct |
| `xml.NewDecoder(r).Token()` | Streaming parser (StartElement, EndElement, CharData) |
| `xml.EscapeText` | Escape text for output |
| `xml.Header` | The `<?xml ...?>` declaration |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./02-io-and-data/xml/
```

## Gotchas
- **Only exported fields are used**, and the element name defaults to the field name.
- **The root element name comes from `XMLName xml.Name` or the type name.** Set it explicitly with an `XMLName` field and tag.
- **`Marshal` does not write the XML declaration.** Prepend `xml.Header` yourself.
- **Namespaces are awkward.** Unmarshal matches by local name unless you give the namespace in the tag.
- **Use the token API (`Decoder.Token`) for large files** instead of loading the tree.
- **Whitespace between elements shows up as `CharData` tokens.** Skip empty ones.
- **Mixed content (text and elements interleaved) is hard to model**; use `xml:",innerxml"` or tokens.
- **Never parse untrusted XML with unbounded size**; wrap the reader in `io.LimitReader`.

## Java comparison

| Java | Go |
|---|---|
| JAXB `@XmlRootElement` | `XMLName xml.Name` field |
| `@XmlAttribute` | `xml:"id,attr"` |
| `@XmlValue` | `xml:",chardata"` |
| StAX `XMLStreamReader` | `xml.Decoder.Token` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/encoding/xml
- Source: https://cs.opensource.google/go/go/+/master:src/encoding/xml/
