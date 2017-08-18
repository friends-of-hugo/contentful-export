package extract

import (
	"../mapper"
	"../read"
	"../translate"
	"../write"

	"log"
)

type Extractor struct {
	ReadConfig read.ReadConfig
	Getter     read.Getter
	Store      write.Store
}

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

	skip := 0

	e.processItem(cf, typeResult, skip)

}
func (e *Extractor) processItem(cf read.Contentful, typeResult mapper.TypeResult, skip int) {
	itemsReader, err := cf.Items(skip)
	if err != nil {
		log.Fatal(err)
	}
	itemResult, err := mapper.MapItems(itemsReader)
	if err != nil {
		log.Fatal(err)
	}
	writer := write.Writer{e.Store}
	tc := translate.TranslationConfig{itemResult}
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
		e.processItem(cf, typeResult, nextPage)
	}
}
