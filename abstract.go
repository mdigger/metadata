package metadata

import (
	"io"

	"gopkg.in/yaml.v3"
)

// Abstract describe default abstract metadata format.
type Abstract []yaml.Node

// Encode encode metadata to YAML.
func (a Abstract) Encode(w io.Writer) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(a)
}
