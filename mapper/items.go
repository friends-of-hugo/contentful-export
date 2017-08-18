package mapper


type ItemResult struct {
	Total int
	Skip  int
	Limit int
	Items []Item
}

type Item struct {
	Sys    Sys
	Fields map[string]interface{}
}


func (item *Item) ContentType() string {
	return item.Sys.ContentType.Sys.ID
}


type Sys struct {
	Type        string
	LinkType    string
	ID          string
	Space       map[string]interface{}
	CreatedAt   string
	Locale      string
	Revision    int
	UpdatedAt   string
	ContentType ContentType
}

type ContentType struct {
	Sys TypeDetails
}

type TypeDetails struct {
	Type     string
	LinkType string
	ID       string
}