package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"io/ioutil"

	"github.com/naoina/toml"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		println("HTTP err")
		return err
	}
	println("HTTP OK ")
	defer r.Body.Close()

	//body, err := ioutil.ReadAll(r.Body)

	//return json.Unmarshal(body, target)
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

type SmallgroupParams struct {
	Title        string
	Slug         string
	Description  string
	LocationText string
	Lat          float64
	Long         float64
	Weekday      string
	Time         string
	//MainContent  string
}

type Smallgroup struct {
	Params      SmallgroupParams
	MainContent string
}

func (s Smallgroup) String() string {
	result := "+++\n"
	output, err := toml.Marshal(s.Params)
	if err != nil {
		return "ERR"
	}

	result += string(output)
	result += "+++\n"
	result += s.MainContent

	err = ioutil.WriteFile("./smallgroups/"+s.Params.Slug+".md", []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

type Sys struct {
	Type      string
	LinkType  string
	ID        string
	Space     map[string]interface{}
	CreatedAt string
	Locale    string
	Revision  int
	UpdatedAt string
}

func main() {

	//var foo2 map[string]interface{}
	var foo2 Foo
	err := getJson("https://cdn.contentful.com/spaces/fp8h0eoshqd0/entries?access_token=2fd06acb06dc3314b28cbd3428be4a3fa9ba2163530f71a09e49ae4c11462006&limit=200&content_type=smallgroup", &foo2)
	if err != nil {
		println("ERR")
		log.Fatal(err)
	}
	println("OK")

	//fmt.Printf("Results: %v\n", foo2)

	for _, item := range foo2.Items {
		//output, err := toml.Marshal(convertSmallgroup(item.Fields))
		//Map := item.Fields
		//output, err := toml.Marshal(Map["title"])
		//if err != nil {
		//	log.Fatal(err)
		//}
		output := convertSmallgroup(item.Fields)
		fmt.Printf("****\n\n%s\n\n****", output)
	}

}

func convertSmallgroup(Map map[string]interface{}) Smallgroup {
	return Smallgroup{
		SmallgroupParams{
			Map["title"].(string),
			Map["slug"].(string),
			Map["description"].(string),
			Map["locationText"].(string),
			Map["locationCoordinates"].(map[string]interface{})["lat"].(float64),
			Map["locationCoordinates"].(map[string]interface{})["lon"].(float64),
			Map["weekday"].(string),
			Map["time"].(string),
		},
		Map["mainContent"].(string),
	}
}
