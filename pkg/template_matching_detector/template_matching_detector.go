package templateMatchingDetector

import (
	"image"
	"image/color"
	_ "image/jpeg"

	cv "gocv.io/x/gocv"
)

func Detect(filename string, template, img cv.Mat) {
	// Ищем лицо на изображении
	result := cv.NewMat()
	cv.MatchTemplate(img, template, &result, cv.TmSqdiffNormed, cv.NewMat())

	_, _, topLeft, _ := cv.MinMaxLoc(result)

	bottomRight := image.Point{X: topLeft.X + template.Cols(), Y: topLeft.Y + template.Rows()}
	cv.Rectangle(&img, image.Rectangle{Min: topLeft, Max: bottomRight}, color.RGBA{A: 255, R: 255, G: 255, B: 255}, 1)

	cv.IMWrite(filename+"_output.jpg", img)
}
