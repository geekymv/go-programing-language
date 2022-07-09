package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDirV3(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	entries := direntsV3(dir)
	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			// 并发读取每个目录
			go walkDirV3(subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func direntsV3(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	return entries
}

func printDiskUsageV3(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1fMB\n", nfiles, float64(nbytes)/(1024*1024))
}

func main() {
	v := flag.Bool("v", false, "show process")

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		// 并发读取每个目录下的文件
		go walkDirV3(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64

	var tick <-chan time.Time
	if *v {
		tick = time.Tick(2 * time.Millisecond)
	}

	for {
		select {
		case <-tick:
			printDiskUsageV3(nfiles, nbytes)
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
	printDiskUsageV3(nfiles, nbytes)
}
