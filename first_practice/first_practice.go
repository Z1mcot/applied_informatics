package first_practice

import (
	"strconv"

	"github.com/z1mcot/applied_informatics/internal"
	templateDetector "github.com/z1mcot/applied_informatics/pkg/template_matching_detector"
	vj "github.com/z1mcot/applied_informatics/pkg/viola_jones_detector"
)

func RunApp() {
	templates, imgs := internal.LoadAssets()

	for templateIndex, template := range templates {
		for imageIndex, img := range imgs {
			templateDetector.Detect(strconv.Itoa(templateIndex)+"_"+strconv.Itoa(imageIndex), template, img)
		}
	}

	models := vj.GetModels()
	for _, model := range models {
		for imageIndex, img := range imgs {
			vj.Detect(strconv.Itoa(imageIndex), model, img)
		}
	}
}
