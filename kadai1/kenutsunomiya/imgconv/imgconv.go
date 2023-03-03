package imgconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var supportExtensions = []string{"jpeg", "jpg", "png"}

type ImageConverter struct {
	From string
	To   string
}

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
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func (ic *ImageConverter) IsValidConv() bool {
	res := map[string]bool{"from": false, "to": false}
	for _, ext := range supportExtensions {
		if ext == ic.From {
			res[ic.From] = true
		}
		if ext == ic.To {
			res[ic.To] = true
		}
	}
	return res[ic.From] && res[ic.To]
}
