package metadata

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Version subscribe the publication version string.
type Version string

var reVersion = regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){2,}$`)

// UnmarshalYAML implement yaml.Unmarshaler interface.
func (ver *Version) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}

	// check version format
	if err := checkVersionFormat(v); err != nil {
		return err
	}

	*ver = Version(v)
	return nil
}

// MarshalYAML implement yaml.Marshaler interface.
func (ver Version) MarshalYAML() (interface{}, error) {
	var v = string(ver)
	if err := checkVersionFormat(v); err != nil {
		return nil, err
	}

	return v, nil
}

func checkVersionFormat(v string) error {
	if !reVersion.MatchString(v) {
		return fmt.Errorf("bad version %q", v)
	}
	return nil
}
