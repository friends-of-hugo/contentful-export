package translate

import "github.com/bhsi-cinch/contentful-hugo/mapper"

// EstablishDirLevelConf provides the ability to augment content directories with with LeafBundle (index.md)
// or Section level (_index.md) frontmatter during the export process.
func EstablishDirLevelConf(t mapper.Type, tc TransConfig) (string, string) {
	var fileName string
	var content string
	if tc.Section[t.Sys.ID] != nil {
		fileName = SectionFilename(t)
		if tc.Encoding == "yaml" {
			content = WriteYamlFrontmatter(tc.Section[t.Sys.ID])
		} else {
			content = WriteTomlFrontmatter(tc.Section[t.Sys.ID])
		}
	}
	if tc.LeafBundle[t.Sys.ID] != nil {
		fileName = LeafBundleFilename(t)
		if tc.Encoding == "yaml" {
			content = WriteYamlFrontmatter(tc.LeafBundle[t.Sys.ID])
		} else {
			content = WriteTomlFrontmatter(tc.LeafBundle[t.Sys.ID])
		}
	}

	return fileName, content
}
