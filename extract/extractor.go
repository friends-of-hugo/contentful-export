package extract

import (
	"github.com/3DHubs/contentful-hugo-export/mapper"
	"github.com/3DHubs/contentful-hugo-export/read"
	"github.com/3DHubs/contentful-hugo-export/translate"
	"github.com/3DHubs/contentful-hugo-export/write"

	"log"
)

// Extractor enables the automated tests to replace key functionalities
// with fakes, mocks and stubs by parameterizing the Reader Configuration,
// the HTTP Getter and the File Store.
type Extractor struct {
	ReadConfig  read.ReadConfig
	Getter      read.Getter
	RStore      read.Store
	TransConfig translate.TransConfig
	WStore      write.Store
}

// ProcessAll goes through all stages: Read, Map, Translate and Write.
// Underwater, it uses private function processItems to allow reading
// through multiple pages of items being returned from Contentful.
func (e *Extractor) ProcessAll() error {

	cf := read.Contentful{
		Getter:     e.Getter,
		ReadConfig: e.ReadConfig,
	}
	typesReader, err := cf.Types()
	if err != nil {
		log.Fatal(err)
		return err
	}

	typeResult, err := mapper.MapTypes(typesReader)
	if err != nil {
		log.Fatal(err)
		return err
	}

	writer := write.Writer{Store: e.WStore}
	for _, t := range typeResult.Items {
		fileName, content := translate.EstablishDirLevelConf(t, e.TransConfig)
		if fileName != "" && content != "" {
			writer.SaveToFile(fileName, content)
		}
	}

	skip := 0

	var ids map[string]string
	ids = make(map[string]string)
	ids = e.processContentTypes(cf, typeResult, skip, ids)

	e.processItems(cf, typeResult, skip, ids)
	return nil
}

// processContentTypes is a recursive function which goes through all pages
// to make a maping of ids to contentTypes
func (e *Extractor) processContentTypes(cf read.Contentful, typeResult mapper.TypeResult, skip int, ids map[string]string) (map[string]string ) {
	itemsReader, err := cf.Items(skip)
	if err != nil {
	    log.Fatal(err)
	}
	itemResult, err := mapper.MapItems(itemsReader)
	if err != nil {
	    log.Fatal(err)
	}

	for _, item := range itemResult.Items {
	    contentType := item.ContentType()
	    ids[item.Sys.ID] = contentType
	}

	nextPage := itemResult.Skip + itemResult.Limit
	if nextPage < itemResult.Total {
	    var nextPageIds = e.processContentTypes(cf, typeResult, nextPage, ids)
	    for k, v := range nextPageIds {
	    ids[k] = v
	    }
	}
	return ids
}


// processItems is a recursive function which goes through all pages
// returned by Contentful and creates a markdownfile for each.
func (e *Extractor) processItems(cf read.Contentful, typeResult mapper.TypeResult, skip int, ids map[string]string) {
	itemsReader, err := cf.Items(skip)
	if err != nil {
		log.Fatal(err)
	}

	itemResult, err := mapper.MapItems(itemsReader)
	if err != nil {
		log.Fatal(err)
	}

	archetypeDataMap := make(map[string]map[string]interface{})
	reader := read.Reader{Store: e.RStore}
	writer := write.Writer{Store: e.WStore}
	tc := translate.TranslationContext{Result: itemResult, TransConfig: e.TransConfig}
	for _, item := range itemResult.Items {
		contentType := item.ContentType()
		itemType, err := typeResult.GetType(contentType)
		if err != nil {
			log.Fatalln(err)
		}

		if archetypeDataMap[contentType] == nil {
			result, err := reader.ViewFromFile(translate.GetArchetypeFilename(contentType))
			if err == nil {
				archeMap, err := tc.TranslateFromMarkdown(result)
				if err != nil {
					log.Fatalln(err)
				}

				archetypeDataMap[contentType] = archeMap
			} else {

				archetypeDataMap[contentType] = make(map[string]interface{})
			}
		}

		contentMap := tc.MapContentValuesToTypeNames(item.Fields, itemType.Fields, ids)
		overriddenContentmap := tc.MergeMaps(archetypeDataMap[contentType], contentMap)
		contentMarkdown := tc.TranslateToMarkdown(tc.ConvertToContent(overriddenContentmap))
		fileName := translate.Filename(item)
		writer.SaveToFile(fileName, contentMarkdown)
	}

	nextPage := itemResult.Skip + itemResult.Limit
	if nextPage < itemResult.Total {
		e.processItems(cf, typeResult, nextPage, ids)
	}
}
