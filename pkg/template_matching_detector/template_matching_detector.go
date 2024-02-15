package templateMatchingDetector

import (
	"image"
	_ "image/jpeg"

	"github.com/z1mcot/applied_informatics/internal"
	"github.com/z1mcot/applied_informatics/internal/colors"
	"gocv.io/x/gocv"
)

func Detect(filename string, template, img gocv.Mat) *internal.ImageForSave {
	// Ищем лицо на изображении
	result := gocv.NewMat()
	gocv.MatchTemplate(img, template, &result, gocv.TmSqdiffNormed, gocv.NewMat())

	_, _, topLeft, _ := gocv.MinMaxLoc(result)

	bottomRight := image.Point{X: topLeft.X + template.Cols(), Y: topLeft.Y + template.Rows()}
	imgCopy := img.Clone()
	gocv.Rectangle(&imgCopy, image.Rectangle{Min: topLeft, Max: bottomRight}, colors.Red, 1)

	filename = "template_matching/" + filename + "_output.jpg"
	return &internal.ImageForSave{Path: filename, Image: imgCopy}
}
