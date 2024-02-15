package vjdetector

import (
	"log"

	"github.com/z1mcot/applied_informatics/internal"
	"github.com/z1mcot/applied_informatics/internal/colors"
	"gocv.io/x/gocv"
)

type VJModel string

const (
	FrontFace VJModel = "../assets/haarcascades/haarcascade_frontalface_default.xml"
	// Eyes              = "../assets/haarcascades/haarcascade_eye_tree_eyeglasses.xml"
)

func GetModels() [1]VJModel { return [1]VJModel{FrontFace} }

var modelNames = map[VJModel]string{
	FrontFace: "front_face",
	// Eyes:      "eyes",
}

func Detect(filename string, model VJModel, img gocv.Mat) *internal.ImageForSave {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(string(model)) {
		log.Fatal("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return nil
	}

	imgCopy := img.Clone()

	detectedAreas := classifier.DetectMultiScale(img)
	for _, r := range detectedAreas {
		gocv.Rectangle(&imgCopy, r, colors.Red, 1)
	}

	filename = "viola_jones/" + filename + "_" + modelNames[model] + ".jpg"

	return &internal.ImageForSave{Path: filename, Image: imgCopy, DetectedAreas: detectedAreas}
}
