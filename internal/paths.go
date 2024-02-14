package internal

import (
	"os"
)

type Path string

func Join[T string | Path](parts ...T) string {
	res := ""
	for _, pathPart := range parts {
		res += "/" + string(pathPart)
	}
	return res
}

const (
	assetsPath Path = "../assets"
	Templates       = assetsPath + "/templates"
	Images          = assetsPath + "/images"
)

func GetAssetsDir() Path {
	return assetsPath
}

func GetAssets(assetsDir Path) ([]os.DirEntry, error) {
	return os.ReadDir(string(assetsDir))
}
