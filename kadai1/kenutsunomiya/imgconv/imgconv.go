package imgconv

var supportExtensions = []string{"jpeg", "jpg", "png"}

type ImageConverter struct {
	From string
	To   string
}

func (ic *ImageConverter) Convert(path string) error {
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
