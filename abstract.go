package metadata

import (
	"io"

	"gopkg.in/yaml.v3"
)

// Abstract is a default not specified metadata format.
type Abstract yaml.Node

// IsZero return true if metadata not defined.
func (a Abstract) IsZero() bool {
	return len(a.Content) == 0
}

// Encode encode metadata to YAML.
func (a Abstract) Encode(w io.Writer) error {
	if a.IsZero() {
		return nil
	}
	return Encode(w, a.Content[0])
}
