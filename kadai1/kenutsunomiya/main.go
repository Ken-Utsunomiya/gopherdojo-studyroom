package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/projects/gopherdojo-studyroom/kadai1/kenutsunomiya/imgconv"
)

var (
	dirFlag  string
	fromFlag string
	toFlag   string
	ic       *imgconv.ImageConverter
)

func init() {
	flag.StringVar(&dirFlag, "dir_name", "", "absolute directory name where you have pictures that you want to convert format")
	flag.StringVar(&fromFlag, "from", "jpeg", "picture format that you want to convert from")
	flag.StringVar(&toFlag, "to", "png", "picture format that you want to convert to")
}

func main() {
	flag.Parse()

	ic = &imgconv.ImageConverter{
		From: fromFlag,
		To:   toFlag,
	}

	if !ic.IsValidConv() {
		log.Fatalf("Error: from=%s, to=%s is not supported conversion\n", ic.From, ic.To)
	}

	dirPath := fmt.Sprintf("./%s", dirFlag)
	if err := walkDir(dirPath); err != nil {
		log.Fatalf("%v\n", err)
	}
}

// traverses all of the files under root directory
func walkDir(root string) error {
	return filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// convert image format
		if !info.IsDir() && filepath.Ext(path)[1:] == ic.From {
			return ic.Convert(path)
		}
		return nil
	})
}
