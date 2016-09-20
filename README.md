# Metadata (YAML front matter)

[![GoDoc](https://godoc.org/github.com/mdigger/metadata?status.svg)](https://godoc.org/github.com/mdigger/metadata)
[![Build Status](https://travis-ci.org/mdigger/metadata.svg)](https://travis-ci.org/mdigger/metadata)

As already in several projects I had to repeatedly describe the work with the metadata headers in the file (YAML front matter), we decided to make it as a separate library. It is still not complete, but I'm slowly copy it duplicate pieces of functionality.

In principle, all built around the simple concept that the metadata describes the class:

```go
type Metadata map[string]interface{}
```

All the rest — the nuances around this. For details, see the source code, as from time to time they are added or changed, and to follow every time according to what is written in the README, not very good.

In short: this is a set of helper functions to download and parse the metadata. Well, obtaining from them the desired values in a convenient form with a minimal amount of code.
