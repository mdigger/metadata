package metadata

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestPublicationConverter(t *testing.T) {
	data := `---
title:
- type: main
  text: My Book
  file-as: book, my
- type: subtitle
  text: An investigation of metadata
creator:
- role: author
  text: John Smith
- role: editor
  text: Sarah Jones
identifier:
- scheme: URN
  text: urn:uuid:02B1F386-E83A-4454-B6EC-422DD949BE43
publisher:  My Press
rights: © 2007 John Smith, CC BY-NC
date: 2021-01
...`

	meta, err := Parse([]byte(data))
	if err != nil {
		t.Fatal(err)
	}

	epub := meta.EPUB()

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	err = enc.Encode(epub)
	if err != nil {
		t.Fatal(err)
	}
}
