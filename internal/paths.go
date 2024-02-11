package internal

import (
	"os"
)

var assetsPath string = "/Users/richardzubko/Projects/applied_informatics/assets"
var isInitialised bool = false

func GetAssetsDir() string {
	return assetsPath
}

func GetAssets() ([]os.DirEntry, error) {
	return os.ReadDir(assetsPath)
}
