package main

import (
	"flag"
	"os"

	"github.com/icyitscold/contentful-hugo/extract"
	"github.com/icyitscold/contentful-hugo/read"
	"github.com/icyitscold/contentful-hugo/translate"
	"github.com/icyitscold/contentful-hugo/write"
)

func main() {
	extractor := extract.Extractor{
		read.ReadConfig{
			"https://cdn.contentful.com",
			*flag.String("space-id", os.Getenv("CONTENTFUL_API_SPACE"), "The contentful space id to export data from"),
			*flag.String("api-key", os.Getenv("CONTENTFUL_API_KEY"), "The contentful delivery API access token"),
			"en-US",
		},
		read.HttpGetter{},
		translate.LoadConfig(),
		write.FileStore{},
	}

	extractor.ProcessAll()
}
