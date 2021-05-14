package metadata

import (
	"encoding/hex"
	"fmt"
	"strings"

	epub "github.com/mdigger/epub3"
	"gopkg.in/yaml.v3"
)

// Identifier of publication.
type Identifier struct {
	Scheme string `yaml:",omitempty"`
	Text   string
}

type idType Identifier // alias

// MarshalYAML implement yaml.Marshaler interface.
func (id Identifier) MarshalYAML() (interface{}, error) {
	if id.Scheme == "" || id.Scheme == "UUID" {
		// if !strings.HasPrefix(id.Text, "urn:uuid:") {
		// 	return fmt.Sprintf("urn:uuid:%v", id.Text), nil
		// }
		return id.Text, nil // only id
	}
	return (idType)(id), nil // as is
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (id *Identifier) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*id = Identifier{Text: value.Value}
	case yaml.MappingNode:
		if err := value.Decode((*idType)(id)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported identifier type: %v", value.Kind)
	}
	if id.Scheme == "" {
		if strings.HasPrefix(id.Text, "urn:uuid:") {
			id.Scheme = "UUID"
			// (*id).Text = strings.TrimPrefix(id.Text, "urn:uuid:")
		} else if strings.HasPrefix(id.Text, "urn:isbn:") {
			id.Scheme = "ISBN"
			// (*id).Text = strings.TrimPrefix(id.Text, "urn:isbn:")
		} else if strings.HasPrefix(id.Text, "doi:") {
			id.Scheme = "DOI"
			// (*id).Text = strings.TrimPrefix(id.Text, "doi:")
		} else {
			// try as uuid
			id.Scheme = "UUID"
			// check uid format
			var text = id.Text
			for _, byteGroup := range []int{8, 4, 4, 4, 12} {
				if text[0] == '-' {
					text = text[1:]
				}
				if _, err := hex.DecodeString(text[:byteGroup]); err != nil {
					id.Scheme = ""
					break
				}
				text = text[byteGroup:]
			}
		}
	}
	return nil
}

// Identifiers describe array of Identifier
type Identifiers []Identifier

// MarshalYAML implement yaml.Marshaler interface.
func (ids Identifiers) MarshalYAML() (interface{}, error) {
	switch l := len(ids); l {
	case 0:
		return &yaml.Node{
			Kind:        yaml.ScalarNode,
			LineComment: epub.NewUUID(),
		}, nil
	case 1:
		return ids[0], nil
	default:
		return []Identifier(ids), nil
	}
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (ids *Identifiers) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		var id Identifier
		if err := value.Decode(&id); err != nil {
			return err
		}
		*ids = Identifiers{id}
	case yaml.SequenceNode:
		*ids = make(Identifiers, len(value.Content))
		for i, node := range value.Content {
			if err := node.Decode(&(*ids)[i]); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported identifiers type: %v", value.Kind)
	}
	return nil
}
