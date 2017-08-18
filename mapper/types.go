package mapper

import "errors"

type TypeResult struct {
	Total int
	Skip  int
	Limit int
	Items []Type
}

func (t *TypeResult) GetType(name string) (result Type, err error) {
	for _, el := range t.Items {
		if el.Sys.ID == name {
			return el, nil
		}
	}
	return Type{}, errors.New("Type not found")
}

type Type struct {
	Sys    Sys
	Name   string
	Fields []TypeField
}

type TypeField struct {
	ID        string
	Name      string
	Type      string
	Localized bool
	Required  bool
	Disabled  bool
	Omitted   bool
}
