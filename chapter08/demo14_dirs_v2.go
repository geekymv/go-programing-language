package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func walkDirV2(dir string, fileSizes chan<- int64) {
	entries := direntsV2(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDirV2(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func direntsV2(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	return entries
}

func printDiskUsageV2(nfiles, nbytes int64) {
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
			walkDirV2(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64

	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-tick:
			printDiskUsageV2(nfiles, nbytes)
		case size, ok := <-fileSizes:
			// 判断 channel 是否关闭
			if !ok {
				fmt.Println("exit")
				goto print
			}
			nfiles++
			nbytes += size
		}
	}

print:
	printDiskUsageV2(nfiles, nbytes)
}
