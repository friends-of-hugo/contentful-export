package translate

import "github.com/bhsi-cinch/contentful-hugo/mapper"

func EstablishSection(t mapper.Type, tc TransConfig) (string, string) {
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
	return fileName, content
}
