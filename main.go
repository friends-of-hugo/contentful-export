package main

import (
	"os"

	"./extract"
	"./read"
	"./translate"
	"./write"
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
