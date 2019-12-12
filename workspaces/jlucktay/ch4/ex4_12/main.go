package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Search terms: %s\n", os.Args[1:])

	wd, errWD := os.Getwd()
	if errWD != nil {
		panic(errWD)
	}

	xkcdJSONPath := filepath.Join(wd, "xkcd")
	files, err := ioutil.ReadDir(xkcdJSONPath)

	if err != nil {
		panic(err)
	}

	transcripts := make(map[int]string)

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			fmt.Printf("'%s' is not a JSON file, skipping!", file.Name())
			continue
		}

		jsonFilePath := filepath.Join(xkcdJSONPath, file.Name())
		content, err := ioutil.ReadFile(jsonFilePath)

		if err != nil {
			fmt.Println("File:", file.Name())
			panic(err)
		}

		comic := new(XKCDComic)
		if errUm := json.Unmarshal(content, &comic); errUm != nil {
			fmt.Println("Content:", content)
			panic(errUm)
		}

		transcripts[comic.Num] = comic.Transcript
	}

	fmt.Printf("# of transcripts: %#v\n", len(transcripts))

	for num, transcript := range transcripts {
		for _, needle := range os.Args[1:] {
			lTranscript := strings.ToLower(transcript)
			lNeedle := strings.ToLower(needle)

			if strings.Contains(lTranscript, lNeedle) {
				fmt.Printf("https://xkcd.com/%d/\n", num)
			}
		}
	}
}

type XKCDComic struct {
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
}
