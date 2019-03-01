// Use the `html/template` package (§4.6) to replace `printTracks` with a function
// that displays the tracks as an HTML table. Use the solution to the previous
// exercise to arrange that each click on a column head makes an HTTP request to
// sort the table.
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"andr.io/ch7/ex7_9/music"
)

var homeDir = "/Users/betandr/workspace/study-group/workspaces/betandr/src/andr.io/ch7/ex7_9"

var order = []music.Attribute{music.Title, music.Artist, music.Album, music.Length, music.Year}

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
		// todo get params or use default
		renderTracks(w)
	}

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(homeDir+"/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir(homeDir+"/fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir(homeDir+"/js"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir(homeDir+"/images"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor", http.FileServer(http.Dir(homeDir+"/vendor"))))
	http.HandleFunc("/", handler)

	http.ListenAndServe(":1337", nil)
}

func renderTracks(out io.Writer) {
	filePrefix, _ := filepath.Abs("./src/andr.io/ch7/ex7_9/templates/")
	tracksList := template.Must(template.ParseFiles(filePrefix + "/index.html"))

	pl := new(music.Playlist)
	pl.Tracks = tracks

	pl.OrderBy(order)

	type Out struct {
		Tracks      []*music.Track
		TitleParam  string
		ArtistParam string
		AlbumParam  string
		YearParam   string
		LengthParam string
	}

	outData := Out{
		Tracks:      pl.Tracks,
		TitleParam:  "title,params",
		ArtistParam: "artist,params",
		AlbumParam:  "album,params",
		YearParam:   "year,params",
		LengthParam: "length,params",
	}

	if err := tracksList.Execute(out, outData); err != nil {
		log.Fatal(err)
	}
}
