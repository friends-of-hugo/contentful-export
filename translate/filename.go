package translate

import (
	"strings"

	"github.com/bhsi-cinch/contentful-hugo/mapper"
)

const baseArchetypeDir string = "./archetypes/"
const baseContentDir string = "./content/"
const idxFile string = "index.md"
const sectionIdxFile string = "_" + idxFile

func Dir(baseDir string, contentType string) string {
	dir := baseDir
	if contentType != "homepage" {
		dir += strings.ToLower(contentType) + "/"
	}

	return dir
}

func Filename(item mapper.Item) string {
	baseDir := baseContentDir
	dir := Dir(baseDir, item.ContentType())
	if dir == baseDir {
		return dir + sectionIdxFile
	}

	return dir + item.Sys.ID + ".md"
}

func SectionFilename(t mapper.Type) string {
	dir := Dir(baseContentDir, t.Sys.ID)

	return dir + sectionIdxFile
}

func LeafBundleFilename(t mapper.Type) string {
	dir := Dir(baseContentDir, t.Sys.ID)

	return dir + idxFile
}

// ArcheTypeFilename takes a content-type's name and returns the file path to the corresponding archetype file.
func GetArchetypeFilename(contentTypeName string) string {
	dir := baseArchetypeDir

	return dir + contentTypeName + ".md"
}
