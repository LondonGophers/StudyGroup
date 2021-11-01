// Write a version of `du` that compute and periodically displays separate totals
// for each of the `root` directories.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// periodic update timer
	var tick <-chan time.Time
	tick = time.Tick(1000 * time.Millisecond)

	// channel for each root
	rootChans := make(map[string](chan int64))

	var n sync.WaitGroup

	// total filesizes
	fileSizes := make(chan int64)

	for _, root := range roots {
		n.Add(1)
		rootChans[root] = make(chan int64)
		go walkDir(root, &n, rootChans[root])
	}

	go func() {
		n.Wait()
		for _, root := range roots {
			close(rootChans[root])
		}
		close(fileSizes)
	}()

	for k, c := range rootChans {
		go listen(k, c, tick, fileSizes)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // chan was closed
			}
			nfiles++
			nbytes += size
		}
	}

	printDiskUsage("total", nfiles, nbytes) // final totals
}

// listen loops over each channel for updates
func listen(name string, c chan int64, tick <-chan time.Time, fileSizes chan<- int64) {
	var nfiles, nbytes int64
loop:
	//!+3
	for {
		select {
		case size, ok := <-c:
			if !ok {
				break loop // chan was closed
			}
			nfiles++
			nbytes += size
			fileSizes <- nbytes
		case <-tick:
			printDiskUsage(name, nfiles, nbytes)
		}
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, rootChan chan<- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, rootChan)
		} else {
			rootChan <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	}
	defer func() { <-sema }() // release token

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}

func printDiskUsage(name string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files  %.1f GB\n", name, nfiles, float64(nbytes)/1e9)
}
