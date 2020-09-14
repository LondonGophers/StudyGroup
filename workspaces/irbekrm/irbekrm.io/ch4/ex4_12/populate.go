package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func populate() error {
	fmt.Println("Retrieving data and writing to xkcd.json..")
	index := make(map[string][]comic)
	resp, err := http.Get(currentUrl)
	if err != nil {
		return fmt.Errorf("Error sending http request: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received http status %v. Exiting\n", resp.StatusCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading http response body: %v\n", err)
	}
	c := comic{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return fmt.Errorf("Error unmarshaling json: %v\n", err)
	}
	for i := 1; i <= c.Num; i++ {
		if i == 404 {
			continue
		}
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Error sending http request: %v\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Received http status %v. Exiting\n", resp.StatusCode)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error reading http response body: %v\n", err)
		}
		c := comic{}
		err = json.Unmarshal(data, &c)
		if err != nil {
			return fmt.Errorf("Error unmarshaling json: %v\n", err)
		}
		re := regexp.MustCompile(`[a-zA-Z-']{4,}`)
		tags := re.FindAll([]byte(c.Transcript), -1)
		// fmt.Printf("Transcript is '%#v'\n", c.Transcript)
		// tags := strings.Split(c.Transcript, " ")
		// fmt.Printf("Tags is '%v'\n", tags)

		for _, t := range tags {
			tt := string(t)
			index[tt] = append(index[tt], c)
		}
	}
	jsonData, err := json.Marshal(index)
	if err != nil {
		return fmt.Errorf("Error marshaling index: %v\n", err)
	}
	err = ioutil.WriteFile("xkcd.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %v\n", err)
	}
	// fmt.Println(index)
	return nil
}
