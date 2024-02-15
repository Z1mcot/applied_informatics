package first_practice

import (
	"strconv"

	"github.com/z1mcot/applied_informatics/internal"
	faceSymmetryDetector "github.com/z1mcot/applied_informatics/pkg/face_symmetry_detector"
	templateDetector "github.com/z1mcot/applied_informatics/pkg/template_matching_detector"
	vj "github.com/z1mcot/applied_informatics/pkg/viola_jones_detector"
)

func RunApp() {
	templates, imgs := internal.LoadAssets()

	var templateMatchingRes []*internal.ImageForSave

	for templateIndex, template := range templates {
		for imageIndex, img := range imgs {
			processedImg := templateDetector.Detect(strconv.Itoa(templateIndex)+"_"+strconv.Itoa(imageIndex), template, img)
			templateMatchingRes = append(templateMatchingRes, processedImg)
		}
	}

	internal.SaveImages(templateMatchingRes)

	models := vj.GetModels()
	var vjRes []*internal.ImageForSave
	for _, model := range models {
		for imageIndex, img := range imgs {
			processedImg := vj.Detect(strconv.Itoa(imageIndex), model, img)
			vjRes = append(vjRes, processedImg)
		}
	}

	internal.SaveImages(vjRes)

	var symRes []*internal.ImageForSave
	for faceIndex, face := range vjRes {
		centralSymmetry := faceSymmetryDetector.DetectSymmetry("", face.Image, face.DetectedAreas, 10)

		for symIndex, symFace := range centralSymmetry {
			localBounds := faceSymmetryDetector.FindLocalSymmetryBounds(symFace.Area, symFace.SymmetryAxis)

			localSym := faceSymmetryDetector.DetectSymmetry(strconv.Itoa(faceIndex)+"_"+strconv.Itoa(symIndex), symFace.Image, localBounds, 10)

			for _, ls := range localSym {
				symRes = append(symRes, &ls.ImageForSave)
			}
		}
	}

	internal.SaveImages(symRes)
}
