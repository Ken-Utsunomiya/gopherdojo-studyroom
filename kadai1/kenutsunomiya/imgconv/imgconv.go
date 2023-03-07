package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var supportExtensions = []string{"jpeg", "jpg", "png", "gif"}

type ImageConverter struct {
	From string
	To   string
}

// converts the file format at the given path
func (ic *ImageConverter) Convert(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	newPath := strings.TrimSuffix(path, ic.From) + ic.To
	new, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer new.Close()

	switch ic.To {
	case "jpg":
	case "jpeg":
		if err := jpeg.Encode(new, img, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(new, img); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(new, img, nil); err != nil {
			return err
		}
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// checks if flags, from and to, are in both supported extensions
func (ic *ImageConverter) IsValidConv() bool {
	res := map[string]bool{"from": false, "to": false}
	for _, ext := range supportExtensions {
		if ext == ic.From {
			res[ic.From] = true
		}
		if ext == ic.To {
			res[ic.To] = true
		}
		if res[ic.From] && res[ic.To] {
			return true
		}
	}
	return false
}
