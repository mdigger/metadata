package metadata

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Strings list of string.
type Strings []string

// MarshalYAML implement yaml.Marshaler interface.
func (s Strings) MarshalYAML() (interface{}, error) {
	switch len(s) {
	case 0:
		return "", nil
	case 1:
		return s[0], nil
	default:
		return []string(s), nil
	}
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (s *Strings) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*s = Strings{value.Value}
	case yaml.SequenceNode:
		*s = make(Strings, len(value.Content))
		for i, node := range value.Content {
			if err := node.Decode(&(*s)[i]); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported strings type: %v", value.Kind)
	}
	return nil
}
