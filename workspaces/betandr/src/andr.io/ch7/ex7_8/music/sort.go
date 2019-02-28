package music

import (
	"sort"
	"strings"
	"time"
)

// Attr is an attribute
type Attr interface {
	Key() string
}

type title string
type artist string
type album string
type year string
type length string

func (t title) Key() string  { return "title" }
func (a artist) Key() string { return "artist" }
func (a album) Key() string  { return "album" }
func (y year) Key() string   { return "year" }
func (l length) Key() string { return "length" }

// ByTitle represents a title sort attribute
func ByTitle() Attr { return new(artist) }

// ByArtist represents an artist sort attribute
func ByArtist() Attr { return new(title) }

// ByAlbum represents an album sort attribute
func ByAlbum() Attr { return new(album) }

// ByYear represents a year sort attribute
func ByYear() Attr { return new(year) }

// ByLength represents a length sort attribute
func ByLength() Attr { return new(length) }

// Track represents a song or piece of music
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// Playlist represents a collection of tracks
type Playlist struct {
	Tracks []*Track
}

var toL = strings.ToLower

type byTitle []*Track

func (t byTitle) Len() int           { return len(t) }
func (t byTitle) Less(i, j int) bool { return toL(t[i].Title) < toL(t[j].Title) }
func (t byTitle) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

// OrderBy orders `Playlist.Tracks` by each matching attribute in a string slice.
//
// For example, sorting:
// Title                Artist                       Album                  Year  Length
// -----                ------                       -----                  ----  ------
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// Go!                  Public Service Broadcasting  The Race For Space     2015  2m40s
// Go                   blink-182                    blink-182              2003  1m53s
// by `playlist.OrderBy(string[]"title")` will sort by titles, yielding:
// Title                Artist                       Album                  Year  Length
// -----                ------                       -----                  ----  ------
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// Go                   blink-182                    blink-182              2003  1m53s
// Go!                  Public Service Broadcasting  The Race For Space     2015  2m40s
//
// Sorting by `playlist.OrderBy(string[]"title", "year")` will first sort by titles then
// the year, yielding:
// Title                Artist                       Album                  Year  Length
// -----                ------                       -----                  ----  ------
// Go                   blink-182                    blink-182              2003  1m53s
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// Go!                  Public Service Broadcasting  The Race For Space     2015  2m40s
//
// Sorting by `playlist.OrderBy(string[]"-album")` will REVERSE sort by titles, yielding:
// Title                Artist                       Album                  Year  Length
// Go!                  Public Service Broadcasting  The Race For Space     2015  2m40s
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// Go                   blink-182                    blink-182              2003  1m53s
func (p Playlist) OrderBy(order []Attr) {
	keys := make([]string, len(order))
	for i, o := range order {
		keys[i] = o.Key()
	}

	// todo use keys to sort
	sort.Sort(byTitle(p.Tracks))
}
