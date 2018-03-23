package main

import (
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
			os.Getenv("SPACE_ID"),
			os.Getenv("CONTENTFUL_KEY"),
			"en-US",
		},
		read.HttpGetter{},
		translate.LoadConfig(),
		write.FileStore{},
	}

	extractor.ProcessAll()
}
