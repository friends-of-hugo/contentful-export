package main

import (
	"flag"
	"os"

	"github.com/bhsi-cinch/contentful-hugo/extract"
	"github.com/bhsi-cinch/contentful-hugo/read"
	"github.com/bhsi-cinch/contentful-hugo/translate"
	"github.com/bhsi-cinch/contentful-hugo/write"
)

func main() {
	space := flag.String("space-id", os.Getenv("CONTENTFUL_API_SPACE"), "The contentful space id to export data from")
	key := flag.String("api-key", os.Getenv("CONTENTFUL_API_KEY"), "The contentful delivery API access token")
	config := flag.String("config-file", "extract-config.toml", "Path to the TOML config file to load for export config")

	flag.Parse()

	extractor := extract.Extractor{
		read.ReadConfig{
			"https://cdn.contentful.com",
			*space,
			*key,
			"en-US",
		},
		read.HttpGetter{},
		translate.LoadConfig(*config),
		write.FileStore{},
	}

	extractor.ProcessAll()
}
