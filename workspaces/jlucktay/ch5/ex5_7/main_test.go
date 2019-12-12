package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestProcess(t *testing.T) {
	htmlBytes, errRead := ioutil.ReadFile("testdata/jameslucktaylor.info.html")
	if errRead != nil {
		t.Fatal(errRead)
	}

	reader := bytes.NewReader(htmlBytes)

	processOut, errProcess := catchStdOut(t, process, reader)
	if errProcess != nil {
		t.Fatal(errProcess)
	}

	processReader := strings.NewReader(processOut)

	if _, errParse := html.Parse(processReader); errParse != nil {
		t.Fatal(errParse)
	}
}

func catchStdOut(t *testing.T, parser func(io.Reader) error, input io.Reader) (string, error) {
	realStdout := os.Stdout

	defer func() { os.Stdout = realStdout }()

	r, fakeStdout, errPipeOpen := os.Pipe()
	if errPipeOpen != nil {
		t.Fatal(errPipeOpen)
	}

	os.Stdout = fakeStdout

	if errParse := parser(input); errParse != nil {
		return "", errParse
	}

	// Need to close here, otherwise ReadAll never gets "EOF".
	if errStdoutClose := fakeStdout.Close(); errStdoutClose != nil {
		t.Fatal(errStdoutClose)
	}

	newOutBytes, errRead := ioutil.ReadAll(r)
	if errRead != nil {
		t.Fatal(errRead)
	}

	if errPipeClose := r.Close(); errPipeClose != nil {
		t.Fatal(errPipeClose)
	}

	return string(newOutBytes), nil
}
