package translate

import (
	"github.com/naoina/toml"
)

func (s Content) ToToml() string {
	result := WriteTomlFrontmatter(s.Params)
	result += s.MainContent

	return result
}

func WriteTomlFrontmatter(fm interface{}) string {
	result := "+++\n"
	output, err := toml.Marshal(fm)
	if err != nil {
		return "ERR"
	}

	result += string(output)
	result += "+++\n"

	return result
}
