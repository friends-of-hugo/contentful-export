package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/adriaandejonge/contentful-hugo/extract"
	"github.com/adriaandejonge/contentful-hugo/read"
	"github.com/adriaandejonge/contentful-hugo/translate"
	"github.com/adriaandejonge/contentful-hugo/write"
)

func main() {
	space := flag.String("space-id", os.Getenv("CONTENTFUL_API_SPACE"), "The contentful space id to export data from")
	key := flag.String("api-key", os.Getenv("CONTENTFUL_API_KEY"), "The contentful delivery API access token")
	config := flag.String("config-file", "extract-config.toml", "Path to the TOML config file to load for export config")
	preview := flag.Bool("p", false, "Use contentful's preview API so that draft content is downloaded")
	flag.Parse()

	fmt.Println("Begin contentful export : ", *space)
	extractor := extract.Extractor{
		ReadConfig: read.ReadConfig{
			UsePreview:  *preview,
			SpaceID:     *space,
			AccessToken: *key,
			Locale:      "en-US",
		},
		Getter:      read.HttpGetter{},
		TransConfig: translate.LoadConfig(*config),
		Store:       write.FileStore{},
	}

	err := extractor.ProcessAll()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("finished")
	}
}
