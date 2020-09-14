package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func search() error {
	index := make(map[string][]comic)
	if len(os.Args) < 3 {
		return errors.New("Usage: xkcdgo search WORD")
	}
	w := os.Args[2]
	fmt.Printf("Searching for %q\n..\n", w)
	if _, err := os.Stat("xkcd.json"); err != nil {
		return fmt.Errorf("Error searching for file 'xkcd.json': %v. Do you need to run 'xkcdgo populate'?\n", err)
	}
	data, err := ioutil.ReadFile("xkcd.json")
	if err != nil {
		return fmt.Errorf("Error reading index file: %v\n", err)
	}
	if err = json.Unmarshal(data, &index); err != nil {
		return fmt.Errorf("Error parsing index: %v\n", err)
	}
	c, ok := index[w]
	if !ok {
		fmt.Printf("No matches found for %q\n", w)
		return nil
	}
	for _, v := range c {
		fmt.Printf("xkcd #%d\nurl: %s\ntranscript: %#v\n\n", v.Num, v.Url, v.Transcript)
	}
	return nil
}
