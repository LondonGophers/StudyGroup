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
)

/////////////////////////////
//    !! IN PROGRESS !!    //
/////////////////////////////

// Track represents a song or piece of music
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{Title: "Ready To Go", Artist: "Republica", Album: "Republica", Year: 1996, Length: length("5m01s")},
	{Title: "Go", Artist: "Lizzo", Album: "Lizzobangers", Year: 2014, Length: length("3m45s")},
	{Title: "Go Get It", Artist: "Slowdive", Album: "Slowdive", Year: 2017, Length: length("6m09s")},
	{Title: "You Gotta Go!", Artist: "The Mighty Mighty Bosstones", Album: "A Jackknife to a Swan", Year: 2002, Length: length("2m43s")},
	{Title: "Go Out", Artist: "Blur", Album: "The Magic Whip", Year: 2015, Length: length("4m41s")},
	{Title: "Just Let Go", Artist: "Fischerspooner", Album: "Odyssey", Year: 2005, Length: length("4m15s")},
	{Title: "Go It Alone", Artist: "Beck", Album: "Guero", Year: 2005, Length: length("4m09s")},
	{Title: "I Let Go", Artist: "Eighteen Visions", Album: "Obsession", Year: 2004, Length: length("3m23s")},
	{Title: "Why Let Go", Artist: "Pearl Jam", Album: "Ten", Year: 1991, Length: length("3m21s")},
	{Title: "Go Next!", Artist: "SOIL&\"PIMP\"SESSIONS", Album: "Planet Pimp", Year: 2008, Length: length("4m59s")},
	{Title: "Go!", Artist: "Public Service Broadcasting", Album: "The Race For Space", Year: 2015, Length: length("2m40s")},
	{Title: "Day Go By", Artist: "Karen O", Album: "Crush Songs", Year: 2014, Length: length("2m16s")},
	{Title: "(Come On) Let's Go!", Artist: "Smashing Pumpkins", Album: "Zeitgeist", Year: 2007, Length: length("3m19s")},
	{Title: "Don't Let Go", Artist: "Warpaint", Album: "Heads Up", Year: 2016, Length: length("4h23s")},
	{Title: "World Go 'Round", Artist: "No Doubt", Album: "Tragic Kingdom", Year: 1995, Length: length("4m09s")},
	{Title: "The Go In The Go-For-It", Artist: "Grandaddy", Album: "", Year: 2003, Length: length("2m59s")},
	{Title: "Lyrics to Go", Artist: "A Tribe Called Quest", Album: "Midnight Marauders", Year: 1993, Length: length("4m09s")},
	{Title: "Boys Say Go!", Artist: "Depeche Mode", Album: "Speak & Spell", Year: 1981, Length: length("3m06s")},
	{Title: "Have Another Go", Artist: "Acoustic Ladyland", Album: "Living with a Tiger", Year: 2009, Length: length("3m59s")},
	{Title: "Go With The Flow", Artist: "Queens Of The Stone Age", Album: "Songs For The Deaf", Year: 2002, Length: length("3m08s")},
	{Title: "Go", Artist: "blink-182", Album: "blink-182", Year: 2003, Length: length("1m53s")},
	{Title: "Go", Artist: "Chemical Brothers", Album: "Born In The Echoes", Year: 2015, Length: length("4m20s")},
}

func length(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return dur
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	printTracks(tracks)
}
