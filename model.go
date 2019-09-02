package bandcamp

import "time"

// Media represents media
type Media struct {
	Type        string    `json:"type"`
	ArtworkURL  string    `json:"artwork_url"`
	Title       string    `json:"title"`
	Album       string    `json:"album"`
	Artist      string    `json:"artist"`
	ReleaseDate time.Time `json:"release_date"`
	URL         string    `json:"url"`
}
