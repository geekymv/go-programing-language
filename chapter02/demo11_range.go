package main

import "os"

func Open(name string) error {
	if f, err := os.Open(name); err != nil {
		return err
	} else {
		f.Close()
	}
	return nil
}

func main() {
}
