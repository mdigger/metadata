// Package metadata is a library for parsing and metadata (YAML front matter) in
// the file to create static websites and blogs.
package metadata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Metadata describes the metadata for documents and projects.
type Metadata map[string]interface{}

// Get returns a string representation of the values contained in the metadata
// under the specified name. In that case, if the data does not already exist,
// then an empty string is returned.
func (self Metadata) Get(name string) string {
	if result, ok := self[name]; ok && result != nil {
		return fmt.Sprint(result)
	}
	return ""
}

// GetQuickList returns the value stored under the specified name in a list of
// strings. If there was initially stored is the string list, it will be
// returned. If the value represented as a string, then it will be broken into
// several lines: as the delimiter stands for any symbol that is not a letter,
// a digit, an underscore or a dash. All other cases will be returned an empty
// list.
func (self Metadata) GetQuickList(name string) []string {
	switch data := self[name].(type) {
	case []string:
		return data
	case []interface{}:
		list := make([]string, len(data))
		for i, value := range data {
			list[i] = fmt.Sprint(value)
		}
		return list
	case string:
		return strings.FieldsFunc(data, func(c rune) bool {
			return c != '_' && c != '-' && !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})
	default:
		return nil
	}
}

// The separator used for separating strings into parts.
var reSplitter = regexp.MustCompile(`\s*[;,]\s*`)

// GetList returns a list of strings stored under the specified name. If there
// is stored a string, it is split into separate lines. As delimiters are the
// comma and the semicolon.
func (self Metadata) GetList(name string) []string {
	switch data := self[name].(type) {
	case []string:
		return data
	case []interface{}:
		list := make([]string, len(data))
		for i, value := range data {
			list[i] = fmt.Sprint(value)
		}
		return list
	case string:
		return reSplitter.Split(strings.TrimSpace(data), -1)
	default:
		return nil
	}
}

// SupportedDatetimeFormats contains the list of supported formats date and time
// that used to parse dates.
var SupportedDatetimeFormats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04Z",
	"2006-01-02T15:04",
	"2006-01-02 15:04:05Z",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04Z",
	"2006-01-02 15:04",
	"2006-01-02Z",
	"2006-01-02",
	"2006-01",
	"2006",
}

// GetDate returns the metadata value with the specified name as a date. If you
// are not able to date or converted from the original format, it returns an
// empty date.
func (self Metadata) GetDate(name string) time.Time {
	switch data := self[name].(type) {
	case time.Time:
		return data
	case int64:
		return time.Unix(data, 0)
	case string:
		for _, format := range SupportedDatetimeFormats {
			var loc *time.Location
			if strings.ContainsRune(format, 'Z') {
				loc = time.UTC
			} else {
				loc = time.Local
			}
			if pdate, err := time.ParseInLocation(format, data, loc); err == nil {
				return pdate
			}
		}
		return time.Time{}
	default:
		return time.Time{}
	}
}

// GetBool returns true if value is defined and looks for approval.
func (self Metadata) GetBool(name string) bool {
	switch data := self[name].(type) {
	case bool:
		return data
	case string:
		value, _ := strconv.ParseBool(data)
		return value
	case int:
		return data > 0
	default:
		return false
	}
}

// GetInt returns a numeric value or zero.
func (self Metadata) GetInt(name string) int {
	switch data := self[name].(type) {
	case string:
		value, _ := strconv.Atoi(data)
		return value
	case int:
		return data
	default:
		return 0
	}
}

// GetSubMetadata returns value with the specified key as metadata.
func (self Metadata) GetSubMetadata(name string) Metadata {
	switch data := self[name].(type) {
	case Metadata:
		return data
	case map[interface{}]interface{}:
		metadata := make(Metadata, len(data))
		for key, value := range data {
			metadata[fmt.Sprint(key)] = value
		}
		return metadata
	default:
		return make(Metadata)
	}
}

// Predefined names of metadata fields.
const (
	MetanameTitle       = "title"       // Title
	MetanameSubtitle    = "subtitle"    // Subtitle
	MetanameDescription = "description" // Description
	MetanameKeywords    = "keywords"    // Keywords
	MetanameTags        = "tags"        // Tags
	MetanameCategories  = "categories"  // Category
	MetanameDate        = "date"        // Date
	MetanameAuthor      = "author"      // Author
	MetanameTemplate    = "layout"      // The name of the template
	MetanameLang        = "lang"        // Language
	MetanameDraft       = "draft"       // The flag of the draft
)

// Title returns the title.
func (self Metadata) Title() string {
	return self.Get(MetanameTitle)
}

// Subtitle returns subtitle.
func (self Metadata) Subtitle() string {
	return self.Get(MetanameSubtitle)
}

// Description returns description.
func (self Metadata) Description() string {
	return self.Get(MetanameDescription)
}

// Keywords returns a list of key words.
func (self Metadata) Keywords() []string {
	return self.GetList(MetanameKeywords)
}

// Tags returns the list of tags.
func (self Metadata) Tags() []string {
	return self.GetQuickList(MetanameTags)
}

// Categories returns the list of categories.
func (self Metadata) Categories() []string {
	return self.GetList(MetanameCategories)
}

// Layout returns the name of the template.
func (self Metadata) Layout() string {
	return self.Get(MetanameTemplate)
}

// Authors returns the list of authors.
func (self Metadata) Authors() []string {
	return self.GetList(MetanameAuthor)
}

// Author returns the first author from the authors list.
func (self Metadata) Author() string {
	return self.Authors()[0]
}

// Date returns the date from the meta information.
func (self Metadata) Date() time.Time {
	return self.GetDate(MetanameDate)
}

// Lang returns the language of the document.
func (self Metadata) Lang() string {
	return self.Get(MetanameLang)
}
