package main

import (
	"os"
	"./read"
	"./write"
	"./extract"
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
		write.FileStore{},
	}

	extractor.ProcessAll()
}