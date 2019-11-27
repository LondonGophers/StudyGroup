package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	htmlBytes, errRead := ioutil.ReadFile("testdata/jameslucktaylor.info.html")
	if errRead != nil {
		t.Fatal(errRead)
	}

	reader := bytes.NewReader(htmlBytes)

	doc, err := html.Parse(reader)
	if err != nil {
		t.Fatal(err)
	}

	element := ElementById(doc, "greeting")

	if element.FirstChild == nil || element.FirstChild.Type != html.TextNode {
		t.Error("The 'greeting' element should have at least one child text node")
	}

	if element.FirstChild.Data != "My CV" {
		t.Error("Incorrect greeting found")
	}
}
