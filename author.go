package metadata

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Author of publication.
type Author struct {
	Role   string `yaml:"role,omitempty"`
	Text   string `yaml:"text"`
	FileAs string `yaml:"file-as,omitempty"`
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (author *Author) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*author = Author{Text: value.Value}
	case yaml.MappingNode:
		type tmpType Author
		if err := value.Decode((*tmpType)(author)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported author type: %v", value.Kind)
	}
	return nil
}

// MarkCode return MARC Code string for Author Role.
func (author Author) MARC() string {
	return MARCCodes[strings.ToLower(author.Role)]
}

// Authors is a list of Author.
type Authors []Author

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (authors *Authors) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*authors = Authors{Author{Text: value.Value}}
	case yaml.SequenceNode:
		*authors = make(Authors, len(value.Content))
		for i, node := range value.Content {
			if err := node.Decode(&(*authors)[i]); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported authors type: %v", value.Kind)
	}
	return nil
}

// MarshalYAML implement yaml.Marshaler interface.
func (authors Authors) MarshalYAML() (interface{}, error) {
	switch len(authors) {
	case 0:
		return &yaml.Node{
			Kind:        yaml.ScalarNode,
			LineComment: "Author Name",
		}, nil
	case 1:
		var author = authors[0]
		if author.Role == "" && author.FileAs == "" {
			return author.Text, nil
		}
		return author, nil
	default:
		var list = make([]string, len(authors))
		for i, author := range authors {
			if author.Role != "" || author.FileAs != "" {
				return ([]Author)(authors), nil
			}
			list[i] = author.Text
		}
		return list, nil
	}
}
