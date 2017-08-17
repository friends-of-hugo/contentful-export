package dump

import "errors"

type Type struct {
	Total int
	Skip  int
	Limit int
	Items []TypeItem
}

func (t *Type) GetType(name string) (result TypeItem, err error) {
	for _, el := range t.Items {
		if el.Sys.ID == name {
			return el, nil
		}
	}
	return TypeItem{}, errors.New("Type not found")
}

type TypeItem struct {
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
