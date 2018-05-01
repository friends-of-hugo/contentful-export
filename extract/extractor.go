package extract

import (
	"github.com/bhsi-cinch/contentful-hugo/mapper"
	"github.com/bhsi-cinch/contentful-hugo/read"
	"github.com/bhsi-cinch/contentful-hugo/translate"
	"github.com/bhsi-cinch/contentful-hugo/write"

	"log"
)

// Extractor enables the automated tests to replace key functionalities
// with fakes, mocks and stubs by parameterizing the Reader Configuration,
// the HTTP Getter and the File Store.
type Extractor struct {
	ReadConfig  read.ReadConfig
	Getter      read.Getter
	TransConfig translate.TransConfig
	Store       write.Store
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

	writer := write.Writer{Store: e.Store}
	for _, t := range typeResult.Items {
		fileName, content := translate.EstablishDirLevelConf(t, e.TransConfig)
		if fileName != "" && content != "" {
			writer.SaveToFile(fileName, content)
		}
	}

	skip := 0

	e.processItems(cf, typeResult, skip)
	return nil
}

// processItems is a recursive function going through all pages
// returned by Contentful
func (e *Extractor) processItems(cf read.Contentful, typeResult mapper.TypeResult, skip int) {
	itemsReader, err := cf.Items(skip)
	if err != nil {
		log.Fatal(err)
	}

	itemResult, err := mapper.MapItems(itemsReader)
	if err != nil {
		log.Fatal(err)
	}

	writer := write.Writer{Store: e.Store}
	tc := translate.TranslationContext{Result: itemResult, TransConfig: e.TransConfig}
	for _, item := range itemResult.Items {

		itemType, err := typeResult.GetType(item.ContentType())
		if err != nil {
			log.Fatal(err)
		}
		fileName, content := tc.Translate(item, itemType)

		writer.SaveToFile(fileName, content)
	}

	nextPage := itemResult.Skip + itemResult.Limit
	if nextPage < itemResult.Total {
		e.processItems(cf, typeResult, nextPage)
	}
}
