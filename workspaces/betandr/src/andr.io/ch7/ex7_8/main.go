// Many GUIs provide a table widget with a stateful multi-tier sort: the primary
// sort key is the most recently clicked column head, the secondary sort key is the
// second most clicked column head, and so on. Define an implementation of
// `sort.Interface` for use by such a table. Compare that approach with repeated
// sorting using `sort.Stable`.
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"andr.io/ch7/ex7_8/music"
)

var tracks = []*music.Track{
	{Title: "Go", Artist: "Chemical Brothers", Album: "Born In The Echoes", Year: 2015, Length: length("4m20s")},
	{Title: "(Come On) Let's Go!", Artist: "Smashing Pumpkins", Album: "Zeitgeist", Year: 2007, Length: length("3m19s")},
	{Title: "Ready To Go", Artist: "Republica", Album: "Republica", Year: 1996, Length: length("5m01s")},
	{Title: "Go Big", Artist: "The Mighty Mighty Bosstones", Album: "A Jackknife to a Swan", Year: 2002, Length: length("2m53s")},
	{Title: "Just Let Go", Artist: "Fischerspooner", Album: "Odyssey", Year: 2005, Length: length("4m15s")},
	{Title: "Go!", Artist: "Public Service Broadcasting", Album: "The Race For Space", Year: 2015, Length: length("2m40s")},
	{Title: "Lyrics to Go", Artist: "A Tribe Called Quest", Album: "Midnight Marauders", Year: 1993, Length: length("4m09s")},
	{Title: "Boys Say Go!", Artist: "Depeche Mode", Album: "Speak & Spell", Year: 1981, Length: length("3m06s")},
	{Title: "Go With The Flow", Artist: "Queens Of The Stone Age", Album: "Songs For The Deaf", Year: 2002, Length: length("3m08s")},
	{Title: "Go", Artist: "blink-182", Album: "blink-182", Year: 2003, Length: length("1m53s")},
	{Title: "Go Your Own Way", Artist: "Fleetwood Mac", Album: "Rumours", Year: 1977, Length: length("3m43s")},
}

func length(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return dur
}

func printPlaylist(pl *music.Playlist) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range pl.Tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	pl := new(music.Playlist)
	pl.Tracks = tracks
	order := []music.Attr{music.ByTitle(), music.ByArtist(), music.ByAlbum(), music.ByYear(), music.ByLength()}

	pl.OrderBy(order)
	printPlaylist(pl)
}
