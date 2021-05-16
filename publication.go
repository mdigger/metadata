package metadata

import (
	"fmt"
	"os"
	"strings"

	epub "github.com/mdigger/epub3"
	"gopkg.in/yaml.v3"
)

// Publication metadata.
type Publication struct {
	Identifier          Identifiers `yaml:"identifier"`
	Title               Titles      `yaml:"title"`
	Language            string      `yaml:"lang,omitempty"` // or legacy: language
	Date                Date        `yaml:"date,omitempty"`
	Creator             Authors     `yaml:"creator"`
	Contributor         Authors     `yaml:"contributor,omitempty"`
	Subject             Subject     `yaml:"subject,omitempty,flow"`
	Description         string      `yaml:"description,omitempty"`
	Type                string      `yaml:"type,omitempty"`
	Format              string      `yaml:"format,omitempty"`
	Publisher           string      `yaml:"publisher,omitempty"`
	Source              string      `yaml:"source,omitempty"`
	Relation            string      `yaml:"relation,omitempty"`
	Coverage            string      `yaml:"coverage,omitempty"`
	Rights              string      `yaml:"rights,omitempty"`
	BelongsToCollection string      `yaml:"belongs-to-collection,omitempty"` // identifies the name of a collection to which the EPUB Publication belongs.
	GroupPosition       string      `yaml:"group-position,omitempty"`        // indicates the numeric position in which the EPUB Publication belongs relative to other works belonging to the same belongs-to-collection field.
	CoverImage          string      `yaml:"cover-image,omitempty"`
	Stylesheets         []string    `yaml:"css,omitempty"` // or legacy: stylesheet
	// PageDirection
	IBooks *struct {
		Version        Version `yaml:"version,omitempty"`
		SpecifiedFonts bool    `yaml:"specified-fonts,omitempty"`
	} `yaml:"ibooks,omitempty"`
	Properties map[string]interface{} `yaml:",omitempty,inline"`
}

// Parse return parsed publication metadata.
func Parse(data []byte) (*Publication, error) {
	pub := new(Publication)
	if err := yaml.Unmarshal(data, pub); err != nil {
		return nil, err
	}

	// check lang synonym
	if lang, ok := pub.Properties["language"]; ok {
		if lang, ok := lang.(string); ok && pub.Language == "" {
			pub.Language = lang
		}
		delete(pub.Properties, "language")
	}

	// check css synonym
	if css, ok := pub.Properties["stylesheet"]; ok {
		switch css := css.(type) {
		case string:
			pub.Stylesheets = append(pub.Stylesheets, css)
		case []string:
			pub.Stylesheets = append(pub.Stylesheets, css...)
		default:
			return nil, fmt.Errorf("bad stylesheet value type: %T", css)
		}
		delete(pub.Properties, "stylesheet")
	}

	return pub, nil
}

// Load return parsed publication metadata from file.
func Load(filename string) (*Publication, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// EPUB return converted to EPUB3 Metadata data.
func (p Publication) EPUB() (meta epub.Metadata) {
	meta.DC = "http://purl.org/dc/elements/1.1/" // add namespace

	// generate ID function
	generateID := func(prefix string, position int) string {
		return fmt.Sprintf("%s-%02d", prefix, position+1)
	}

	// identifiers
	for i, identifier := range p.Identifier {
		id := generateID("id", i)
		meta.Identifier = append(meta.Identifier, epub.Element{
			Value: identifier.Text, ID: id})

		if identifier.Scheme != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "identifier-type",
				Scheme:   identifier.Scheme,
				Value:    identifier.Onix(),
			})
		}
	}

	// titles
	for i, title := range p.Title {
		var id string
		if title.FileAs != "" || title.Type != "" {
			id = generateID("title", i)
		}

		meta.Title = append(meta.Title, epub.ElementLang{
			Value: title.Text, ID: id})

		if title.Type != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "title-type",
				Value:    title.Type,
			})
		}

		if title.FileAs != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "file-as",
				Value:    title.FileAs,
			})
		}
	}

	// lang
	if p.Language != "" {
		meta.Language = []epub.Element{{Value: p.Language}}
	}

	// date
	if p.Date != "" {
		meta.Date = &epub.Element{Value: string(p.Date)}
	}

	// creators
	for i, creator := range p.Creator {
		role := creator.MARC()

		var id string
		if creator.FileAs != "" || role != "" {
			id = generateID("creator", i)
		}

		meta.Creator = append(meta.Creator, epub.ElementLang{
			Value: creator.Text, ID: id})

		if role != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "role",
				Scheme:   "marc:relators",
				Value:    role,
			})
		}

		if creator.FileAs != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "file-as",
				Value:    creator.FileAs,
			})
		}
	}

	// contributors
	for i, contributor := range p.Contributor {
		role := contributor.MARC()

		var id string
		if contributor.FileAs != "" || role != "" {
			id = generateID("contributor", i)
		}

		meta.Contributor = append(meta.Contributor, epub.ElementLang{
			Value: contributor.Text, ID: id})

		if role != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "role",
				Scheme:   "marc:relators",
				Value:    role,
			})
		}

		if contributor.FileAs != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "file-as",
				Value:    contributor.FileAs,
			})
		}
	}

	// subjects
	for _, subject := range p.Subject {
		meta.Subject = append(meta.Subject, epub.ElementLang{Value: subject})
	}

	// description
	if p.Description != "" {
		// remove new line & spaces
		descripion := strings.Join(strings.Fields(p.Description), " ")
		meta.Description = []epub.ElementLang{{Value: descripion}}
	}

	// type
	if p.Type != "" {
		meta.Type = []epub.Element{{Value: p.Type}}
	}

	// format
	if p.Format != "" {
		meta.Format = []epub.Element{{Value: p.Format}}
	}

	// publisher
	if p.Publisher != "" {
		meta.Publisher = []epub.ElementLang{{Value: p.Publisher}}
	}

	// source
	if p.Source != "" {
		meta.Source = []epub.Element{{Value: p.Source}}
	}

	// relation
	if p.Relation != "" {
		meta.Relation = []epub.ElementLang{{Value: p.Relation}}
	}

	// coverage
	if p.Coverage != "" {
		meta.Coverage = []epub.ElementLang{{Value: p.Coverage}}
	}

	// rights
	if p.Rights != "" {
		meta.Rights = []epub.ElementLang{{Value: p.Rights}}
	}

	// collection
	if p.BelongsToCollection != "" {
		id := generateID("collection", 0)
		meta.Meta = append(meta.Meta, epub.Meta{
			ID:       id,
			Property: "belongs-to-collection",
			Value:    p.BelongsToCollection,
		})

		if p.GroupPosition != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "collection-type",
				Value:    "series",
			})

			meta.Meta = append(meta.Meta, epub.Meta{
				Refines:  id,
				Property: "group-position",
				Value:    p.GroupPosition,
			})
		}
	}

	// ibooks
	if p.IBooks != nil {
		// version
		if p.IBooks.Version != "" {
			meta.Meta = append(meta.Meta, epub.Meta{
				Property: "ibooks:version",
				Value:    string(p.IBooks.Version),
			})
		}

		// specified fonts
		if p.IBooks.SpecifiedFonts {
			meta.Meta = append(meta.Meta, epub.Meta{
				Property: "ibooks:specified-fonts",
				Value:    "yes",
			})
		}
	}

	return
}
