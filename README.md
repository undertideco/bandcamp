# go-bandcamp
[![GoDoc](https://godoc.org/github.com/undertideco/go-bandcamp?status.svg)](https://godoc.org/github.com/undertideco/go-bandcamp)

This is a library that can do the following on Bandcamp:
- Search
- Lookup Tracks/Albums by URL

### Usage
```go
package main

import (
  "github.com/undertideco/go-bandcamp"
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
