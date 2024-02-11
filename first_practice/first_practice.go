package first_practice

import (
	"strconv"

	"github.com/z1mcot/applied_informatics/internal"
	detector "github.com/z1mcot/applied_informatics/pkg/template_matching_detector"
)

func RunApp() {
	template, imgs := internal.LoadAssetsForTemplateMathching()

	for index, img := range imgs {
		detector.Detect(strconv.Itoa(index), template, img)
	}
}
