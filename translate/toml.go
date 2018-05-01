package translate

import (
	"github.com/naoina/toml"
)

// ToToml .() -> string
// Takes a Content struct and outputs it as TOML frontmatter followed by main-content.
func (s Content) ToToml() string {
	result := WriteTomlFrontmatter(s.Params)
	result += s.MainContent

	return result
}

// WriteTomlFrontmatter (fm Map[]) -> string
// Converts a Map[] into a TOML string, pre and postfixing it with `+++` to designate frontmatter.
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
