package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

var (
	dirFlag  string
	fromFlag string
	toFlag   string
)

func init() {
	flag.StringVar(&dirFlag, "dir_name", "", "absolute directory name where you have pictures that you want to convert format")
	flag.StringVar(&fromFlag, "from", "jpeg", "picture format that you want to convert from")
	flag.StringVar(&toFlag, "to", "png", "picture format that you want to convert to")
}

func main() {
	flag.Parse()

	dirPath := fmt.Sprintf("./%s", dirFlag)
	if err := walkDir(dirPath); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func walkDir(root string) error {
	return filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// convert image format
		fmt.Printf("name: %s\n", info.Name())

		return nil
	})
}
