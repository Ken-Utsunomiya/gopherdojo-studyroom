package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/projects/gopherdojo-studyroom/kadai1/kenutsunomiya/imgconv"
)

func main() {
	fmt.Println("Starting Image Converter...")

	// TODO: read a dir name from the command line
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter dir name: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error: Failed to read dir name.")
			continue
		}
		dirName := strings.TrimSuffix(line, "\n")
		// path, err := os.Getwd()
		if err != nil {
			fmt.Printf("❌ Error: Failed to find the dir with name = %s\n", dirName)
			continue
		}

		path := fmt.Sprintf("./%s", dirName)
		if err = traverseDir(path); err != nil {
			fmt.Printf("❌ Error: Failed to read at %v\n", err)
		} else {
			fmt.Println("✅ Conversion Successful")
		}
		break
	}
}

// TODO: convert all JPG files under the specified dir to PNG files
func traverseDir(dirPath string) error {
	// TODO: find the specified dir
	//       dir not found -> re-prompt an input
	//       dir found     -> recursively traverse files under dir and convert
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if dirErr := traverseDir(path); dirErr != nil {
				return dirErr
			}
		}
		imgBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return imgconv.JpgToPng(imgBytes)
	})
	return err
}
