package metadata

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Delimiter metadata from the data.
var splitMetadataLine = []byte("\n---\n")

// ReadFile reads the file metadata, parses it, and returns with the rest of
// the text.
func ReadFile(filename string) (metadata Metadata, data []byte, err error) {
	metadata = make(Metadata)
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	if parts := bytes.SplitN(data, splitMetadataLine, 2); len(parts) == 2 {
		if err = yaml.Unmarshal(parts[0], metadata); err != nil {
			return
		}
		data = parts[1]
	}
	return
}
