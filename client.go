package bandcamp

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	// MediaTypeAlbum represents an album
	MediaTypeAlbum = "album"

	// MediaTypeTrack represents a track
	MediaTypeTrack = "track"

	timeFormatSearch = "02 January 2006"
	timeFormatLookup = "20060102"
)

var (
	regexAlbum  = regexp.MustCompile(`(?s)from (.+?)\s.+? by`)
	regexArtist = regexp.MustCompile(`by (.+)`)

	regexSearchReleaseDate = regexp.MustCompile(`released (\d{2} \w+ \d{4})`)
)

// Client represents the request client
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client Client) collector() *colly.Collector {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "bandcamp.com/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})
	return c
}

// Search search for music
func (client Client) Search(term string) ([]Media, error) {
	var results []Media
	var err error

	c := client.collector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".searchresult", func(e *colly.HTMLElement) {
		var releaseDate *time.Time
		albumMatches := regexAlbum.FindStringSubmatch(e.ChildText(".subhead"))
		artistMatches := regexArtist.FindStringSubmatch(e.ChildText(".subhead"))
		rawDateMatches := regexSearchReleaseDate.FindStringSubmatch(e.ChildText(".released"))
		if len(rawDateMatches) >= 2 {
			var rd time.Time
			rd, err = time.Parse(timeFormatSearch, rawDateMatches[1])
			if err == nil {
				releaseDate = &rd
			}
		}
		r := Media{
			Type:        strings.ToLower(strings.TrimSpace(e.ChildText(".itemtype"))),
			ArtworkURL:  e.ChildAttr(".artcont img", "src"),
			Title:       e.ChildText(".heading a"),
			URL:         e.ChildText(".itemurl a"),
			ReleaseDate: releaseDate,
		}

		switch r.Type {
		case MediaTypeAlbum:
		case MediaTypeTrack:
			break
		default:
			return
		}

		if len(albumMatches) > 1 {
			r.Type = MediaTypeTrack
			r.Album = albumMatches[1]
		} else {
			r.Type = MediaTypeAlbum
		}
		r.Artist = artistMatches[1]
		results = append(results, r)
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	c.Visit(fmt.Sprintf("https://bandcamp.com/search?q=%s", url.QueryEscape(term)))

	c.Wait()

	return results, err
}

// Lookup look up for a resource
func (client Client) Lookup(url string) (Media, error) {
	var result = Media{
		Type: MediaTypeTrack,
		URL:  url,
	}
	var err error

	c := client.collector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".trackView .trackTitle", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".fromAlbum", func(e *colly.HTMLElement) {
		result.Album = strings.TrimSpace(e.Text)
	})

	c.OnHTML("span[itemprop=byArtist]", func(e *colly.HTMLElement) {
		result.Artist = strings.TrimSpace(e.Text)
	})

	c.OnHTML("img[itemprop=image]", func(e *colly.HTMLElement) {
		result.ArtworkURL = e.Attr("src")
	})

	c.OnHTML("meta[itemprop=datePublished]", func(e *colly.HTMLElement) {
		var releaseDate time.Time
		releaseDate, err = time.Parse(timeFormatLookup, e.Attr("content"))
		if err == nil {
			result.ReleaseDate = &releaseDate
		}
	})

	c.OnHTML(".trackView", func(e *colly.HTMLElement) {
		if e.Attr("itemtype") == "http://schema.org/MusicAlbum" {
			result.Type = MediaTypeAlbum
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
		err = fmt.Errorf("%+v", r.StatusCode)
	})

	c.Visit(url)

	c.Wait()

	return result, err
}
