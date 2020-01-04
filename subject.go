package metadata

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Subject list of publication.
type Subject []string

// MarshalYAML implement yaml.Marshaler interface.
func (subject Subject) MarshalYAML() (interface{}, error) {
	switch len(subject) {
	case 0:
		return "", nil
	case 1:
		return subject[0], nil
	default:
		return []string(subject), nil
	}
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (subject *Subject) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*subject = Subject{value.Value}
	case yaml.SequenceNode:
		*subject = make(Subject, len(value.Content))
		for i, node := range value.Content {
			if err := node.Decode(&(*subject)[i]); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported subject type: %v", value.Kind)
	}
	return nil
}
