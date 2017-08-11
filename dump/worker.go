package dump

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"io/ioutil"

	"github.com/naoina/toml"
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

type Foo struct {
	Total int
	Skip  int
	Limit int
	Items []Item
}

type Item struct {
	Sys    Sys
	Fields map[string]interface{}
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

func Work() {

	var foo2 Foo
	err := getJson("https://cdn.contentful.com/spaces/"+os.Getenv("SPACE_ID")+"/entries?access_token="+os.Getenv("CONTENTFUL_KEY")+"&limit=200&content_type=smallgroup", &foo2)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range foo2.Items {
		dir := "./content/" + item.Sys.ContentType.Sys.ID + "/"
		var fileMode os.FileMode
		fileMode = 0733
		err := os.MkdirAll(dir, fileMode)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(item.Sys.ContentType.Sys.ID)
		output := convertContent(item.Fields)

		err = ioutil.WriteFile(dir+output.Slug+".md", []byte(output.String()), fileMode)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func convertContent(Map map[string]interface{}) Content {
	return Content{
		map[string]interface{}{
			"title":        Map["title"].(string),
			"slug":         Map["slug"].(string),
			"description":  Map["description"].(string),
			"locationText": Map["locationText"].(string),
			"lat":          Map["locationCoordinates"].(map[string]interface{})["lat"].(float64),
			"long":         Map["locationCoordinates"].(map[string]interface{})["lon"].(float64),
			"weekday":      Map["weekday"].(string),
			"time":         Map["time"].(string),
		},
		Map["mainContent"].(string),
		Map["slug"].(string),
	}
}
