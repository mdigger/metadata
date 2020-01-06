package metadata

import (
	"io"
	"regexp"

	"gopkg.in/yaml.v3"
)

var reMetadata = regexp.MustCompile(`(?s)\A-{3}\n(.+)\n(?:\.{3}|-{3})\n*`)

// Parse split and parse metadata from markdown.
func Parse(markdown []byte, data interface{}) ([]byte, error) {
	result := reMetadata.FindSubmatch(markdown)
	if result == nil {
		return markdown, nil // no metadata found
	}
	if err := yaml.Unmarshal(result[1], data); err != nil {
		return markdown, err // metadata error
	}
	return markdown[len(result[0]):], nil // extract metadata
}

// Encode write YAML front matter.
func Encode(w io.Writer, data interface{}) error {
	if data == nil {
		return nil
	}
	_, err := io.WriteString(w, "---\n")
	if err != nil {
		return err
	}
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, "---\n")
	return err
}
