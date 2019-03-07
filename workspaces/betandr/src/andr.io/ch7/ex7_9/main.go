// Use the `html/template` package (§4.6) to replace `printTracks` with a function
// that displays the tracks as an HTML table. Use the solution to the previous
// exercise to arrange that each click on a column head makes an HTTP request to
// sort the table.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"andr.io/ch7/ex7_8/music"
)

var homeDir = flag.String("home", "", "The full path to the directory containing css, fonts, images, js, templates, and vendor directories")

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
	{Title: "Blitzkrieg Bop", Artist: "Ramones", Album: "Ramones", Year: 1976, Length: length("2m12s")},
	{Title: "Go Your Own Way", Artist: "Fleetwood Mac", Album: "Rumours", Year: 1977, Length: length("3m43s")},
}

func main() {
	flag.Parse()

	if len(*homeDir) == 0 {
		fmt.Println("error: home directory not specified")
		os.Exit(0)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		order := parseOrderParams(r.Form["order"])

		if len(order) == 0 {
			order = []music.Attribute{music.Artist, music.Album, music.Title, music.Year, music.Length}
		} else if len(order) != 5 {
			// using `http.ResponseWriter` interface:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad Request: %d order params found but require 5", len(order))
			// ...or could use the `http.Error` utility function
			// msg := fmt.Sprintf("Bad Request: %d order params found but require 5", len(order))
			// http.Error(w, msg, http.StatusBadRequest)
			return
		}

		renderTracks(w, tracks, order)
	}

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(*homeDir+"/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir(*homeDir+"/fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir(*homeDir+"/js"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir(*homeDir+"/images"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor", http.FileServer(http.Dir(*homeDir+"/vendor"))))
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8000", nil)
}

// length takes a string such as `"4m09s"` and returns a `time.Duration` representation of it.
func length(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return dur
}

// decodeParam takes a string and returns a `music.Attribute` representation of it.
// If not matching `music.Attribute` is found a `-1` is returned.
func decodeParam(param string) music.Attribute {
	switch param {
	case "title":
		return music.Title
	case "artist":
		return music.Artist
	case "album":
		return music.Album
	case "length":
		return music.Length
	case "year":
		return music.Year
	default:
		return -1 // todo - handle default case with something better than ""
	}
}

// decodeAttr takes a `music.Attribute` enum value and returns a string representation of it.
// If not matching `music.Attribute` is found an empty string is returned.
func decodeAttr(attr music.Attribute) string {
	switch attr {
	case 0:
		return "title"
	case 1:
		return "artist"
	case 2:
		return "album"
	case 3:
		return "year"
	case 4:
		return "length"
	default:
		return "" // todo - handle default case with something better than ""
	}
}

// parseOrderParams takes a slice of strings and returns into a slice of `music.Attribute`s.
func parseOrderParams(params []string) (order []music.Attribute) {
	order = make([]music.Attribute, 0)
	for _, p := range params {
		order = append(order, decodeParam(p))
	}
	return
}

// moveAttrToFront moves a specified `music.Attribute` to the front of a slice of
// `music.Attribute`s. The attribute is removed from its position and all other attributes
// moved up one place and then the attribute is inserted into the first position of the list.
func moveAttrToFront(attrToLead music.Attribute, attrs []music.Attribute) {
	idx := func(al music.Attribute, as []music.Attribute) int {
		for i, a := range as {
			if a == al {
				return i
			}
		}
		return -1
	}(attrToLead, attrs)

	for i := idx; i > 0; i-- {
		attrs[i] = attrs[i-1]
	}

	attrs[0] = attrToLead
}

// attributesToValues converts a list of attributes to a string as a `template.URL`. It also
// moves a single attribute to the front of the list. This is used to generate the URL params
// for each column head meaning that the column head being clicked will order the `attrToLead`
// first.
// e.g. considering the column headers one, two, three, and four. For column one is represented
// as the URL:
// `?order=one&order=two&order=three&order=four`
// For column two by the URL:
// `?order=two&order=one&order=three&order=four`
// For column three by the URL:
// `?order=three&order=one&order=two&order=four`
// For column four by the URL:
// `?order=four&order=one&order=two&order=three`
//
// These URLs will order the table by the selected column and preserve the ordering of the
// previously clicked heads, or the default.
// This allows the separation of the table rendering and the sort order "stacking".
func attributesToValues(order []music.Attribute, attrToLead music.Attribute) template.URL {
	moveAttrToFront(attrToLead, order)
	v := url.Values{}
	for _, a := range order {
		s := decodeAttr(a)
		if s != "" {
			v.Add("order", s)
		}
	}
	return template.URL(v.Encode())
}

// renderTracks writes a slice of `music.Track`s represented by pointers to an `io.Writer`
// in the order supplied by a slice of `music.Attribute`s.
func renderTracks(out io.Writer, tracks []*music.Track, order []music.Attribute) {
	filePrefix, _ := filepath.Abs("./src/andr.io/ch7/ex7_9/templates/")
	tracksList := template.Must(template.ParseFiles(filePrefix + "/index.html"))

	pl := new(music.Playlist)
	pl.Tracks = tracks

	pl.OrderBy(order)

	var data struct {
		Tracks      []*music.Track
		TitleParam  template.URL
		ArtistParam template.URL
		AlbumParam  template.URL
		YearParam   template.URL
		LengthParam template.URL
	}

	data.Tracks = pl.Tracks
	data.TitleParam = attributesToValues(order, music.Title)
	data.ArtistParam = attributesToValues(order, music.Artist)
	data.AlbumParam = attributesToValues(order, music.Album)
	data.YearParam = attributesToValues(order, music.Year)
	data.LengthParam = attributesToValues(order, music.Length)

	if err := tracksList.Execute(out, data); err != nil {
		log.Fatal(err)
	}
}
