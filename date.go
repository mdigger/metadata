package metadata

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

// Date subscribe the publication date.
//
// A string value in YYYY-MM-DD format. (Only the year is necessary.)
type Date string

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (date *Date) UnmarshalYAML(value *yaml.Node) (err error) {
	var d string
	if err = value.Decode(&d); err != nil {
		return err
	}
	// check data format
	if err := checkDateFormat(d); err != nil {
		return err
	}

	*date = Date(d)
	return nil
}

// MarshalYAML implement yaml.Marshaler interface.
func (date Date) MarshalYAML() (interface{}, error) {
	var d = string(date)
	if err := checkDateFormat(d); err != nil {
		return nil, err
	}
	return d, nil
}

func checkDateFormat(d string) (err error) {
	// check data format
	var dateTime time.Time
	for _, layout := range []string{"2006-01-02", "2006-01", "2006", time.RFC3339} {
		if dateTime, err = time.Parse(layout, d); err == nil {
			break
		}
	}

	if dateTime.IsZero() {
		return fmt.Errorf("bad date %v", d)
	}

	return nil
}
