package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func walkDir(dir string, fileSizes chan<- int64) {
	entries := dirents(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1fMB\n", nfiles, float64(nbytes)/(1024*1024))
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	printDiskUsage(nfiles, nbytes)
}
