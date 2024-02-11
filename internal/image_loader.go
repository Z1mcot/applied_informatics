package internal

import (
	"image"
	"log"
	"os"
	"strings"

	cv "gocv.io/x/gocv"
)

var extensions [2]string = [2]string{"png", "jpg"}

func LoadAssetsForTemplateMathching() (cv.Mat, []cv.Mat) {
	var template cv.Mat
	var imgs []cv.Mat

	assets, _ := GetAssets()

	for _, element := range assets {
		splitName := strings.Split(element.Name(), ".")
		if splitName[1] != "jpg" && splitName[1] != "jpeg" {
			continue
		}

		path := GetAssetsDir() + "/" + element.Name()
		img := loadImage(path)
		parsedImg, err := cv.ImageToMatRGBA(img)

		if err != nil {
			log.Fatal(err)
		}

		if splitName[0] == "template" {
			template = parsedImg
		} else {
			imgs = append(imgs, parsedImg)
		}
	}

	return template, imgs
}

// Функция для загрузки изображения из файла
func loadImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
