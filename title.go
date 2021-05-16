package metadata

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Title of publication.
//
// Valid values for type are main, subtitle, short, collection, edition, extended.
type Title struct {
	Type   string `yaml:",omitempty"`
	Text   string `yaml:"text"`
	FileAs string `yaml:"file-as,omitempty"`
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (title *Title) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*title = Title{Text: value.Value}
	case yaml.MappingNode:
		type tmpType Title
		if err := value.Decode((*tmpType)(title)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported title type: %v", value.Kind)
	}
	if title.Type == "" {
		title.Type = "main"
	}
	return nil
}

// Titles is a lis of Title.
type Titles []Title

// MarshalYAML implement yaml.Marshaler interface.
func (titles Titles) MarshalYAML() (interface{}, error) {
	switch len(titles) {
	case 0:
		return &yaml.Node{
			Kind:        yaml.ScalarNode,
			LineComment: "Title",
		}, nil
	case 1:
		var title = titles[0]
		if (title.Type == "" || title.Type == "main") && title.FileAs == "" {
			return title.Text, nil
		}
		return title, nil
	default:
		return []Title(titles), nil
	}
}

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (titles *Titles) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*titles = Titles{Title{Type: "main", Text: value.Value}}
	case yaml.SequenceNode: // list
		*titles = make(Titles, len(value.Content))
		for i, node := range value.Content {
			if err := node.Decode(&(*titles)[i]); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported titles type: %v", value.Kind)
	}
	return nil
}

// title return string with title of type tt.
func (titles Titles) title(tt string) string {
	for _, title := range titles {
		if title.Type == tt {
			return title.Text
		}
	}
	return ""
}

// Main return first main title.
func (titles Titles) Main() string {
	return titles.title("main")
}

// Subtitle return first subtitle.
func (titles Titles) Subtitle() string {
	return titles.title("subtitle")
}
