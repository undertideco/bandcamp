# bandcamp
[![Build Status](https://travis-ci.org/undertideco/bandcamp.svg?branch=master)](https://travis-ci.org/undertideco/bandcamp) [![GoDoc](https://godoc.org/github.com/undertideco/bandcamp?status.svg)](https://godoc.org/github.com/undertideco/bandcamp)

This is a library that can do the following on Bandcamp:
- Search
- Lookup Tracks/Albums by URL

### Usage
```go
package main

import (
  "github.com/undertideco/bandcamp"
)

func main() {
  bandcampClient := bandcamp.NewClient()

  results, err := bandcampClient.Search("Coldplay")
  if err != nil {
    log.Println(err)
  }

  track, err := bandcampClient.Lookup("https://spunoutofcontrol.bandcamp.com/track/infinite")
  if err != nil {
    log.Println(err)
  }

  album, err := bandcampClient.Lookup("https://coldplay.bandcamp.com/album/greatest-hits")
  if err != nil {
    log.Println(err)
  }
}
```
