package internal

import (
	"image"

	"gocv.io/x/gocv"
)

type ImageForSave struct {
	Path          string
	Image         gocv.Mat
	DetectedAreas []image.Rectangle
}

func SaveImages(imgs []*ImageForSave) {
	for _, img := range imgs {
		if img == nil {
			continue
		}
		gocv.IMWrite(img.Path, img.Image.Clone())
	}
}
