package first_practice

import (
	"strconv"

	"github.com/z1mcot/applied_informatics/internal"
	templateDetector "github.com/z1mcot/applied_informatics/pkg/template_matching_detector"
)

func RunApp() {
	templates, imgs := internal.LoadAssets()

	for index, img := range imgs {
		templateDetector.Detect(strconv.Itoa(index), templates, img)
	}
}
