package translate

import (
	"github.com/naoina/toml"
)

func (s Content) String() string {
	result := "+++\n"
	output, err := toml.Marshal(s.Params)
	if err != nil {
		return "ERR"
	}

	result += string(output)
	result += "+++\n"
	result += s.MainContent

	return result
}
