package metadata

import (
	"regexp"

	"gopkg.in/yaml.v3"
)

var reMetadata = regexp.MustCompile(`(?s)\A-{3}\n(.+)\n(?:\.{3}|-{3})\n*`)

// Abstract describe default abstract metadata format.
type Abstract = []yaml.Node

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
