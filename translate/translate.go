package translate

import (
	"../mapper"
)

type Content struct {
	Params      map[string]interface{}
	MainContent string
	Slug        string
}

func Translate(item mapper.Item, itemType mapper.Type) (fileName string, content string) {

	content = convertContent(item.Fields, itemType.Fields).String()
	fileName = Filename(item)
	return
}

func convertContent(Map map[string]interface{}, fields []mapper.TypeField) Content {
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

func translateField(value interface{}, field mapper.TypeField) interface{} {
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
