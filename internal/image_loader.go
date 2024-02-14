package internal

import (
	"image"
	"log"
	"os"
	"strings"

	"gocv.io/x/gocv"
)

var extensions = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
}

func LoadAssets() (templates []gocv.Mat, images []gocv.Mat) {
	templateAssets, _ := GetAssets(Templates)
	imgAssets, _ := GetAssets(Images)

	return loadImagesFromPaths(Templates, templateAssets), loadImagesFromPaths(Images, imgAssets)
}

func loadImagesFromPaths(baseDir Path, assets []os.DirEntry) []gocv.Mat {
	images := make([]gocv.Mat, 0, len(assets))
	for _, element := range assets {
		splitName := strings.Split(element.Name(), ".")
		if !extensions[splitName[1]] {
			continue
		}

		path := string(baseDir) + "/" + element.Name()
		parsedImg, err := gocv.ImageToMatRGBA(loadImage(path))

		if err != nil {
			log.Fatal(err)
		}

		images = append(images, parsedImg)
	}
	return images
}

func loadImage(filename string) image.Image {
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatal(openErr)

	}
	defer func(file *os.File) {
		fileErr := file.Close()
		if fileErr != nil {
			log.Fatal(fileErr)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
