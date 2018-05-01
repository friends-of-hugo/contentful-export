package translate

import "gopkg.in/yaml.v2"

// ToYaml .() -> string
// Takes a Content struct and outputs it as YAML frontmatter followed by main-content.
func (s Content) ToYaml() string {
	result := WriteYamlFrontmatter(s.Params)
	result += s.MainContent

	return result
}

// WriteYamlFrontmatter (fm Map[]) -> string
// Converts a Map[] into a YAML string, pre and postfixing it with `---` to designate frontmatter.
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
