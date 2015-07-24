package metadata

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Разделитель метаданных от данных.
var splitMetadataLine = []byte("\n---\n")

// ReadFile читает файл с метаданными, разбирает его и возвращает вместе с оставшимся текстом.
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
