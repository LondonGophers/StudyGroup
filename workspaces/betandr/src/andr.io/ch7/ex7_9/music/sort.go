package music

import (
	"sort"
	"strings"
	"time"
)

type Attribute int

const (
	Title Attribute = iota
	Artist
	Album
	Year
	Length
)

// Heading returns the string representation for a particular attribute
func Heading(attr Attribute) string {
	switch attr {
	case 0:
		return "Title"
	case 1:
		return "Artist"
	case 2:
		return "Album"
	case 3:
		return "Year"
	case 4:
		return "Length"
	default:
		return "Unknown"
	}
}

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

var sortOrder = []Attribute{Title, Artist, Album, Year, Length}

type bySortOrder []*Track

func (t bySortOrder) Len() int      { return len(t) }
func (t bySortOrder) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t bySortOrder) Less(i, j int) bool {
	// uses anonymous functions in a map which can be accessed in the order of sortOrder
	var title func(lt bool) bool
	title = func(lt bool) bool {
		if lt {
			return toL(t[i].Title) < toL(t[j].Title)
		}
		return toL(t[i].Title) > toL(t[j].Title)
	}

	var artist func(lt bool) bool
	artist = func(lt bool) bool {
		if lt {
			return toL(t[i].Artist) < toL(t[j].Artist)
		}
		return toL(t[i].Artist) > toL(t[j].Artist)
	}

	var album func(lt bool) bool
	album = func(lt bool) bool {
		if lt {
			return toL(t[i].Album) < toL(t[j].Album)
		}
		return toL(t[i].Album) > toL(t[j].Album)
	}

	var year func(lt bool) bool
	year = func(lt bool) bool {
		if lt {
			return t[i].Year < t[j].Year
		}
		return t[i].Year > t[j].Year
	}

	var length func(lt bool) bool
	length = func(lt bool) bool {
		if lt {
			return t[i].Length < t[j].Length
		}
		return t[i].Length > t[j].Length
	}

	sorts := make(map[Attribute]func(bool) bool)
	sorts[Title] = title
	sorts[Artist] = artist
	sorts[Album] = album
	sorts[Year] = year
	sorts[Length] = length

	// range over the sort functions in the order they appear in sortOrder
	for _, order := range sortOrder {
		f := sorts[order]
		// `if f(true)` and `if f(false)` isn't the same as `return f(true)`
		// as the bool causes a different code path to be executed in the
		// anonymous function; true = less than comparison, false = greater
		// than comparison. So both `f(true)` _and_ `f(false)` can return
		// `true` or `false` depending on the logical comparison.
		//
		// todo Think about this structure a little more.
		if f(true) {
			return true
		}
		if f(false) {
			return false
		}
	}

	return false
}

// OrderBy sorts `Playlist.Tracks` into an ascending order (small to large, a-z...)
// specified by a slice of `Attribute`s such as `[]Attribute{Year, Album, Title}`.
//
// For example, given the table:
// Title                Artist                       Album                  Year  Length
// -----                ------                       -----                  ----  ------
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// (Come On) Let's Go!  Smashing Pumpkins            Zeitgeist              2007  3m19s
// Go                   blink-182                    blink-182              2003  1m53s
//
// Sorting by `playlist.OrderBy([]Attribute{Title})` will sort by titles, yielding:
// Title                Artist                       Album                  Year  Length
// (Come On) Let's Go!  Smashing Pumpkins            Zeitgeist              2007  3m19s
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
// Go                   blink-182                    blink-182              2003  1m53s
//
// Sorting by `playlist.OrderBy([]Attribute{Year, Title})` will first sort by the year
// then the title, yielding:
// Title                Artist                       Album                  Year  Length
// -----                ------                       -----                  ----  ------
// Go                   blink-182                    blink-182              2003  1m53s
// (Come On) Let's Go!  Smashing Pumpkins            Zeitgeist              2007  3m19s
// Go                   Chemical Brothers            Born In The Echoes     2015  4m20s
func (p Playlist) OrderBy(order []Attribute) {
	if len(order) != 5 {
		panic("must have 5 attributes for this naive implementation")
	}
	// todo use attributes to sort
	sortOrder = order // the sort order is used in `[]*Track.Less`
	sort.Sort(bySortOrder(p.Tracks))
}
