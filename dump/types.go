package dump

type Type struct {
	Total int
	Skip  int
	Limit int
	Items []TypeItem
}

func (t *Type) GetType(name string) TypeItem {
	for _, el := range t.Items {
		if el.Sys.ID == name {
			return el
		}
	}
	// TODO: Throw error - why can't I return nil?
	return t.Items[0]
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
