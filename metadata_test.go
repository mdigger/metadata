package metadata_test

import (
	"fmt"

	"github.com/mdigger/metadata"
	"gopkg.in/yaml.v2"
)

func ExampleMetadata() {
	var data = []byte(`---
title: Title
date: "2014-12-24Z"
list: "test1, test2, test3"
exts: "*.txt;*.bak"
emptylist: []
quicklist: "aa bb cc dd, ee, test2014; bla-bla"
---`)
	var metadata metadata.Metadata
	if err := yaml.Unmarshal(data, &metadata); err != nil {
		panic(err)
	}
	fmt.Println(metadata.Get("title"))
	fmt.Println(metadata.GetDate("date"))
	fmt.Println(metadata.GetList("list"))
	fmt.Println(metadata.GetList("exts"))
	fmt.Println(metadata.GetList("emptylist"))
	fmt.Println(metadata.GetList("none") == nil)
	fmt.Println(metadata.GetQuickList("quicklist"))
	// Output:
	// Title
	// 2014-12-24 00:00:00 +0000 UTC
	// [test1 test2 test3]
	// [*.txt *.bak]
	// []
	// true
	// [aa bb cc dd ee test2014 bla-bla]
}
