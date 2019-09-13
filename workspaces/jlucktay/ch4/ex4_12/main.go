package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	wd, errWD := os.Getwd()
	if errWD != nil {
		panic(errWD)
	}

	xkcdJSONPath := filepath.Join(wd, "xkcd")
	fmt.Printf("walking '%s'...\n", xkcdJSONPath)

	files, err := ioutil.ReadDir(xkcdJSONPath)
	if err != nil {
		panic(err)
	}

	counter := 0

	// transcripts := make(map[int]string)

	for _, file := range files {
		if counter > 10 {
			os.Exit(0)
		}

		jsonFilePath := filepath.Join(xkcdJSONPath, file.Name())

		fmt.Println(jsonFilePath)
		counter++

		content, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("File contents: %s", content)
	}

}
