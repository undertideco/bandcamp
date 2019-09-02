package bandcamp

import (
	"testing"
	"time"
)

var (
	expectedAlbum = Media{
		Type:        MediaTypeAlbum,
		Title:       "Greatest Hits",
		Artist:      "Coldplay",
		ReleaseDate: time.Date(2017, time.October, 28, 0, 0, 0, 0, time.UTC),
	}

	expectedTrack = Media{
		Type:        MediaTypeTrack,
		Title:       "Infinite",
		Album:       "MONOCHROME DAYDREAM",
		Artist:      "BRYCE MILLER",
		ReleaseDate: time.Date(2019, time.November, 1, 0, 0, 0, 0, time.UTC),
	}
)

func TestSearch(t *testing.T) {
	c := NewClient()
	searchData, err := c.Search("chvrches")
	if err != nil {
		t.Error(err)
	}
	if len(searchData) == 0 {
		t.Errorf("Expected at least 1 search result")
	}
}

func TestLookupAlbum(t *testing.T) {
	c := NewClient()
	album, err := c.Lookup("https://coldplay.bandcamp.com/album/greatest-hits")
	if err != nil {
		t.Error(err)
	}
	if album.Type != expectedAlbum.Type {
		t.Errorf("Unexpected album type '%s', expected '%s'", album.Type, expectedAlbum.Type)
	}
	if album.Title != expectedAlbum.Title {
		t.Errorf("Unexpected album title '%s', expected '%s'", album.Title, expectedAlbum.Title)
	}
	if album.Artist != expectedAlbum.Artist {
		t.Errorf("Unexpected album artist '%s', expected '%s'", album.Artist, expectedAlbum.Artist)
	}
	if album.ReleaseDate != expectedAlbum.ReleaseDate {
		t.Errorf("Unexpected album release date '%s', expected '%s'", album.ReleaseDate, expectedAlbum.ReleaseDate)
	}
}

func TestLookupTrack(t *testing.T) {
	c := NewClient()
	track, err := c.Lookup("https://spunoutofcontrol.bandcamp.com/track/infinite")
	if err != nil {
		t.Error(err)
	}
	if track.Type != expectedTrack.Type {
		t.Errorf("Unexpected track type '%s', expected '%s'", track.Type, expectedTrack.Type)
	}
	if track.Title != expectedTrack.Title {
		t.Errorf("Unexpected track title '%s', expected '%s'", track.Title, expectedTrack.Title)
	}
	if track.Artist != expectedTrack.Artist {
		t.Errorf("Unexpected track artist '%s', expected '%s'", track.Artist, expectedTrack.Artist)
	}
	if track.Album != expectedTrack.Album {
		t.Errorf("Unexpected track album '%s', expected '%s'", track.Album, expectedTrack.Album)
	}
	if track.ReleaseDate != expectedTrack.ReleaseDate {
		t.Errorf("Unexpected track release date '%s', expected '%s'", track.ReleaseDate, expectedTrack.ReleaseDate)
	}
}
