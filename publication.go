package metadata

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Publication metadata.
type Publication struct {
	Identifier     Identifiers
	Title          Titles
	Creator        Authors
	Contributor    Authors                `yaml:",omitempty"`
	Publisher      string                 `yaml:",omitempty"`
	Lang           string                 `yaml:",omitempty"`
	Date           Date                   `yaml:",omitempty"`
	Subject        Subject                `yaml:",omitempty,flow"`
	Description    string                 `yaml:",omitempty"`
	Type           string                 `yaml:",omitempty"`
	Format         string                 `yaml:",omitempty"`
	Relation       string                 `yaml:",omitempty"`
	Coverage       string                 `yaml:",omitempty"`
	Rights         string                 `yaml:",omitempty"`
	CoverImage     string                 `yaml:"cover-image,omitempty"`
	CSS            string                 `yaml:",omitempty"`
	Version        string                 `yaml:",omitempty"`
	SpecifiedFonts bool                   `yaml:"specified-fonts,omitempty"`
	Properties     map[string]interface{} `yaml:",omitempty,inline"`
}

// Parse return parsed publication metadata.
func Parse(data []byte) (*Publication, error) {
	pub := new(Publication)
	if err := yaml.Unmarshal(data, pub); err != nil {
		return nil, err
	}
	return pub, nil
}

// Load return parsed publication metadata from file.
func Load(filename string) (*Publication, error) {
	data, err := os.ReadFile("sample.yaml")
	if err != nil {
		return nil, err
	}
	return Parse(data)
}
