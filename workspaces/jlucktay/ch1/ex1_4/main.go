package main

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		log.Fatal("Must provide at least one filename as an argument.")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				log.WithFields(log.Fields{
					"error":    err,
					"filename": arg,
				}).Error("Error opening a file.")
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, m := range counts {
		total := 0
		filenames := make([]string, 0)

		for file, matchCount := range m {
			filenames = append(filenames, file)
			total += matchCount
		}

		if total > 1 {
			log.WithFields(log.Fields{
				"files": strings.Join(filenames, ", "),
				"line":  line,
				"total": total,
			}).Info("Files in which this duplicated line occurs.")
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	s, err := f.Stat()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"file":  f,
		}).Error("Couldn't stat file.")

		return
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		if _, exists := counts[input.Text()]; !exists {
			log.WithFields(log.Fields{
				"key": input.Text(),
			}).Trace("Initialising sub-map within 'counts'.")

			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][s.Name()]++
	}
}
