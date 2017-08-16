package dump

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"io/ioutil"

	"github.com/naoina/toml"
	"strings"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Result struct {
	Total int
	Skip  int
	Limit int
	Items []Item
}

type Item struct {
	Sys    Sys
	Fields map[string]interface{}
}

func (item *Item) Dir() string {
	dir := "./content/"
	contentType := item.ContentType()
	if contentType != "homepage" {
		dir += contentType + "/"
	}
	return dir
}

func (item *Item) Filename() string {
	dir := item.Dir()
	if dir == "./content/" {
		return dir + "_index.md"
	}

	return dir + item.Sys.ID + ".md"
}

func (item *Item) ContentType() string {
	return item.Sys.ContentType.Sys.ID
}

type Content struct {
	Params      map[string]interface{}
	MainContent string
	Slug        string
}

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

type Dumper struct {
	Types Type
	UrlBase string
	SpaceID string
	AccessToken string
	//Config
	// e.g. /content as basedir
	// e.g. mainContent
	// e.g. slug
	// e.g. 200 items per page
	// e.g. homepage -> _index.md
	// etc
}

func (d *Dumper) Work() {
	d.WorkSkip(0)
}
func (d *Dumper) WorkSkip(skip int) {

	var result Result
	err := getJson(d.UrlBase + "/spaces/"+d.SpaceID+"/entries?access_token="+d.AccessToken+"&limit=200&skip="+strconv.Itoa(skip), &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		d.processItem(item)
	}

	nextPage := result.Skip + result.Limit
	if nextPage < result.Total {
		d.WorkSkip(nextPage)
	}
}

func (d *Dumper) processItem(item Item) {
	itemType := d.Types.GetType(item.ContentType())
	output := convertContent(item.Fields, itemType.Fields).String()
	fileName := item.Filename()
	saveToFile(fileName, output)
}

func saveToFile(fileName string, output string) {
	var fileMode os.FileMode
	fileMode = 0733

	err := os.MkdirAll(dirForFile(fileName), fileMode)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, []byte(output), fileMode)
	if err != nil {
		log.Fatal(err)
	}
}

func dirForFile(filename string) string {
	index := strings.LastIndex(filename, "/")
	return filename[0:index]
}


func convertContent(Map map[string]interface{}, fields []TypeField) Content {
	fieldMap := map[string]interface{}{}

	for _, field := range fields {
		fieldMap[field.ID] = translateField(Map[field.ID], field)
	}
	mainContent := removeItem(fieldMap, "mainContent").(string)
	slug, _ := fieldMap["slug"].(string)

	return Content{
		fieldMap,
		mainContent,
		slug,
	}
}

func removeItem(Map map[string]interface{}, toDelete string) interface{} {
	value := Map[toDelete]
	if value == nil {
		return ""
	}
	delete(Map, toDelete)
	return value
}

func translateField(value interface{}, field TypeField) interface{} {
	if field.Type == "Array" {
		items := value.([]interface{})

		var array []string
		array = make([]string, len(items))

		for i, el := range items {
			sys := el.(map[string]interface{})["sys"].(map[string]interface{})
			array[i] = sys["id"].(string) + ".md"
		}
		return array
	}
	return value
}
