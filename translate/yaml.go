package translate

import "gopkg.in/yaml.v2"

func (s Content) ToYaml() string {
	result := WriteYamlFrontmatter(s.Params)
	result += s.MainContent

	return result
}

func WriteYamlFrontmatter(fm interface{}) string {
	result := "---\n"
	output, err := yaml.Marshal(fm)
	if err != nil {
		return "ERR"
	}

	result += string(output)
	result += "---\n"
	return result
}
