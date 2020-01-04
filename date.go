package metadata

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

// Date describe the publication date.
type Date string

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (date *Date) UnmarshalYAML(value *yaml.Node) (err error) {
	var d string
	if err = value.Decode(&d); err != nil {
		return err
	}
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
	*date = Date(d)
	return nil
}
