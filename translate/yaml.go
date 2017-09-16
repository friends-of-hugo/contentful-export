package translate

import "gopkg.in/yaml.v2"

func (s Content) ToYaml() string {
	result := "---\n"
	output, err := yaml.Marshal(s.Params)
	if err != nil {
		return "ERR"
	}

	result += string(output)
	result += "---\n"
	result += s.MainContent

	return result
}
