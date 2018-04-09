package extract

import (
	"github.com/bhsi-cinch/contentful-hugo/mapper"
	"github.com/bhsi-cinch/contentful-hugo/read"
	"github.com/bhsi-cinch/contentful-hugo/translate"
	"github.com/bhsi-cinch/contentful-hugo/write"

	"log"
)

// By parameterizing the Reader Configuration,
// the HTTP Getter and the File Store, it enables the automated tests to
// replace key functionalities with fakes, mocks and stubs.
type Extractor struct {
	ReadConfig  read.ReadConfig
	Getter      read.Getter
	TransConfig translate.TransConfig
	Store       write.Store
}

// ProcessAll goes through all stages: Read, Map, Translate and Write.
// Underwater, it uses private function processItems to allow reading
// through multiple pages of items being returned from Contentful.
func (e *Extractor) ProcessAll() {

	cf := read.Contentful{
		e.Getter,
		e.ReadConfig,
	}
	typesReader, err := cf.Types()
	if err != nil {
		log.Fatal(err)
	}

	typeResult, err := mapper.MapTypes(typesReader)
	if err != nil {
		log.Fatal(err)
	}

	writer := write.Writer{e.Store}
	for _, t := range typeResult.Items {
		fileName, content := translate.EstablishSection(t, e.TransConfig)
		if fileName != "" && content != "" {
			writer.SaveToFile(fileName, content)
		}
	}

	skip := 0

	e.processItems(cf, typeResult, skip)

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
	writer := write.Writer{e.Store}
	tc := translate.TranslationContext{itemResult, e.TransConfig}
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
