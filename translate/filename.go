package translate

import (
	"strings"

	"github.com/bhsi-cinch/contentful-hugo/mapper"
)

const baseDir string = "./content/"
const idxFile string = "index.md"
const rootIdxFile string = "_" + idxFile

func Dir(contentType string) string {
	dir := baseDir
	if contentType != "homepage" {
		dir += strings.ToLower(contentType) + "/"
	}
	return dir
}

func Filename(item mapper.Item) string {
	dir := Dir(item.ContentType())
	if dir == baseDir {
		return dir + rootIdxFile
	}

	return dir + item.Sys.ID + ".md"
}

func SectionFilename(t mapper.Type) string {
	dir := Dir(t.Sys.ID)
	return dir + idxFile
}
