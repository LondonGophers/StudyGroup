package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func search(url string) (*SearchResult, error) {
	resp, errGet := http.Get(url)
	if errGet != nil {
		return nil, errGet
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	buf := new(bytes.Buffer)
	if _, errRead := buf.ReadFrom(resp.Body); errRead != nil {
		return nil, errRead
	}

	var result SearchResult
	if errUm := json.Unmarshal(buf.Bytes(), &result); errUm != nil {
		return nil, errUm
	}

	return &result, nil
}
