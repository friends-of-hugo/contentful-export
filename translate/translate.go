package translate

import (
	"fmt"
	"regexp"

	"github.com/bhsi-cinch/contentful-hugo/mapper"
)

type TranslationContext struct {
	Result      mapper.ItemResult
	TransConfig TransConfig
}

type Content struct {
	Params      map[string]interface{}
	MainContent string
	Slug        string
}

func (tc *TranslationContext) Translate(item mapper.Item, itemType mapper.Type) (fileName string, content string) {

	rawContent := tc.convertContent(item.Fields, itemType.Fields)
	if tc.TransConfig.Encoding == "yaml" {
		content = rawContent.ToYaml()
	} else {
		content = rawContent.ToToml()
	}

	fileName = Filename(item)
	return
}

func (tc *TranslationContext) convertContent(Map map[string]interface{}, fields []mapper.TypeField) Content {
	fieldMap := map[string]interface{}{}

	for _, field := range fields {
		value := tc.translateField(Map[field.ID], field)
		if value != nil {
			fieldMap[field.ID] = value
		}
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

func (tc *TranslationContext) translateField(value interface{}, field mapper.TypeField) interface{} {
	if field.Type == "Array" {
		if value == nil {
			return []interface{}{}
		}
		items := value.([]interface{})

		var array []string
		array = make([]string, len(items))

		for i, el := range items {
			s, isString := el.(string)
			if isString {
				array[i] = s
			} else {
				sys := el.(map[string]interface{})["sys"].(map[string]interface{})
				array[i] = tc.translateLink(sys)
			}
		}
		return array
	} else if field.Type == "Link" {
		if value == nil {
			return value
		}
		item := value.(map[string]interface{})
		sys := item["sys"].(map[string]interface{})

		return tc.translateLink(sys)

	} else if field.Type == "Date" {
		re, err := regexp.Compile(`([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2})(\+[0-9]{2}:[0-9]{2})?`) // want to know what is in front of 'at'
		if err != nil {
			fmt.Println(err)
		}
		res := re.FindAllStringSubmatch(value.(string), -1)
		if len(res) > 0 {
			value = fmt.Sprintf("%v:00%v", res[0][1], res[0][2])
		}

	}
	return value
}

func (tc *TranslationContext) translateLink(sys map[string]interface{}) string {
	linkType := sys["linkType"]
	if linkType == "Entry" {
		return sys["id"].(string) + ".md"
	} else {
		assets := tc.Result.Includes["Asset"]
		for _, asset := range assets {

			if sys["id"].(string) == asset.Sys.ID {
				return asset.Fields["file"].(map[string]interface{})["url"].(string)
			}
		}
		// Look up asset - but from where???
	}
	return "ERR"
}
