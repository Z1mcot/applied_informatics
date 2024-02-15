package faceSymmetryDetector

import (
	"image"
	"strconv"

	"github.com/z1mcot/applied_informatics/internal"
	"github.com/z1mcot/applied_informatics/internal/colors"
	"gocv.io/x/gocv"
)

func splitByLines(input *gocv.Mat, w, symmetryAxis int) (l, r gocv.Mat) {
	flippedL := gocv.NewMat()
	flippedR := gocv.NewMat()

	l = input.Region(image.Rect(symmetryAxis-w, 0, symmetryAxis, input.Rows()))
	gocv.Flip(l, &flippedL, 1)

	r = input.Region(image.Rect(symmetryAxis, 0, symmetryAxis+w, input.Rows()))
	gocv.Flip(l, &flippedR, 1)
	return l, r
}

func calculateDistance(l, r gocv.Mat, w int) float64 {
	resMat := gocv.NewMat()
	gocv.AbsDiff(l, r, &resMat)
	return resMat.Sum().Val1
}

func findMinIndex(arr []float64) int {
	minIndex := 0
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return minIndex
}

type SymmetricImageForSave struct {
	internal.ImageForSave
	SymmetryAxis int
	Area         image.Rectangle
}

func DetectSymmetry(filename string, img gocv.Mat, faceBoundries []image.Rectangle, w int) []*SymmetricImageForSave {
	res := make([]*SymmetricImageForSave, len(faceBoundries))
	imgCopy := img.Clone()
	for i, faceBound := range faceBoundries { // Проходим по всем найденным областям

		if !(0 <= faceBound.Min.X && 0 <= faceBound.Dx() && faceBound.Max.X <= imgCopy.Cols() && 0 <= faceBound.Min.Y && 0 <= faceBound.Dy() && faceBound.Max.Y <= imgCopy.Rows()) {
			return nil
		}

		face := imgCopy.Region(faceBound) // Отделяем только лицо

		var distances []float64

		x1 := 0
		x2 := face.Cols()

		symmetryAxis := x1 + w

		// вместо while
		for symmetryAxis >= (x1+w) && symmetryAxis <= (x2-w) {
			l, r := splitByLines(&face, w, symmetryAxis)
			distance := calculateDistance(l, r, w)
			distances = append(distances, distance)
			symmetryAxis += w
		}

		if len(distances) == 0 {
			return nil
		}

		symmetryAxis = findMinIndex(distances) * w

		pt1 := image.Point{X: symmetryAxis, Y: 0}
		pt2 := image.Point{X: symmetryAxis, Y: face.Rows()}
		gocv.Line(&face, pt1, pt2, colors.Red, 1)

		if filename != "" {
			filename = "face_symmetry_detector/" + strconv.Itoa(i) + "_" + filename + "_output.jpg"
		}
		imgForSave := internal.ImageForSave{Path: filename, Image: imgCopy}
		res[i] = &SymmetricImageForSave{ImageForSave: imgForSave, SymmetryAxis: symmetryAxis, Area: faceBound}
	}

	return res
}

func FindLocalSymmetryBounds(faceArea image.Rectangle, symmetryLine int) []image.Rectangle {
	localBound := make([]image.Rectangle, 2)
	minPoint := image.Point{faceArea.Min.X + faceArea.Dx()/4, faceArea.Min.Y + faceArea.Dy()/4}
	maxPoint := image.Point{faceArea.Min.X + symmetryLine, faceArea.Max.Y - faceArea.Dy()/4}
	localBound[0] = image.Rectangle{Min: minPoint, Max: maxPoint}

	minPoint = image.Point{faceArea.Min.X + symmetryLine, faceArea.Min.Y + faceArea.Dy()/4}
	maxPoint = image.Point{faceArea.Max.X - faceArea.Dx()/4, faceArea.Max.Y - faceArea.Dy()/4}
	localBound[1] = image.Rectangle{Min: minPoint, Max: maxPoint}

	return localBound
}
