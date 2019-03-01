// Use the `html/template` package (§4.6) to replace `printTracks` with a function
// that displays the tracks as an HTML table. Use the solution to the previous
// exercise to arrange that each click on a column head makes an HTTP request to
// sort the table.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"andr.io/ch7/ex7_9/music"
)

var homeDir = "/Users/betandr/workspace/study-group/workspaces/betandr/src/andr.io/ch7/ex7_9"

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

func length(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return dur
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		order := parseOrderParams(r.Form["order"])

		if len(order) == 0 {
			order = []music.Attribute{music.Artist, music.Album, music.Title, music.Year, music.Length}
		} else if len(order) != 5 {
			http.Error(w, fmt.Sprintf("Bad Request: %d order params, require 5", len(order)), http.StatusBadRequest)
			return
		}

		renderTracks(w, order)
	}

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(homeDir+"/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir(homeDir+"/fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir(homeDir+"/js"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir(homeDir+"/images"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor", http.FileServer(http.Dir(homeDir+"/vendor"))))
	http.HandleFunc("/", handler)

	http.ListenAndServe(":1337", nil)
}

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
		return -1 // todo - handle default case with an "unknown"
	}
}

func decodeAttr(attr music.Attribute) string {
	switch attr {
	case 0:
		return "title"
	case 1:
		return "artist"
	case 2:
		return "album"
	case 3:
		return "length"
	case 4:
		return "year"
	default:
		return ""
	}
}

func parseOrderParams(params []string) (order []music.Attribute) {
	order = make([]music.Attribute, 0)
	for _, p := range params {
		order = append(order, decodeParam(p))
	}
	return
}

func moveAttrToFront(attrToLead music.Attribute, attrs *[]music.Attribute) {
	idx := func(al music.Attribute, as *[]music.Attribute) int {
		for i, a := range *as {
			if a == al {
				return i
			}
		}
		return -1
	}(attrToLead, attrs)

	for i := idx; i >= 0; i-- {
		fmt.Printf("move %d to %d\n", i-1, i)
		// todo
	}

	//todo set first item to attrToLead
}

func attributesToValues(order []music.Attribute, attrToLead music.Attribute) template.URL {
	moveAttrToFront(attrToLead, &order)
	v := url.Values{}
	for _, a := range order {
		s := decodeAttr(a)
		if s != "" {
			v.Add("order", s)
		}
	}
	return template.URL(v.Encode())
}

func renderTracks(out io.Writer, order []music.Attribute) {
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
