package translate

import "../mapper"

type TranslationConfig struct {
	Result mapper.ItemResult
}

type Content struct {
	Params      map[string]interface{}
	MainContent string
	Slug        string
}

func (tc *TranslationConfig) Translate(item mapper.Item, itemType mapper.Type) (fileName string, content string) {

	content = tc.convertContent(item.Fields, itemType.Fields).String()
	fileName = Filename(item)
	return
}

func (tc *TranslationConfig) convertContent(Map map[string]interface{}, fields []mapper.TypeField) Content {
	fieldMap := map[string]interface{}{}

	for _, field := range fields {
		fieldMap[field.ID] = tc.translateField(Map[field.ID], field)
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

func (tc *TranslationConfig) translateField(value interface{}, field mapper.TypeField) interface{} {
	if field.Type == "Array" {
		items := value.([]interface{})

		var array []string
		array = make([]string, len(items))

		for i, el := range items {
			sys := el.(map[string]interface{})["sys"].(map[string]interface{})
			array[i] = tc.translateLink(sys)
		}
		return array
	} else if field.Type == "Link" {
		item := value.(map[string]interface{})
		sys := item["sys"].(map[string]interface{})

		return tc.translateLink(sys)

	}
	return value
}

func (tc *TranslationConfig) translateLink(sys map[string]interface{}) string {
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
