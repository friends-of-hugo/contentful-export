package main

import (
	"os"

	"./dump"
	"./impl"
)

func main() {

	dumper := dump.Dumper{
		dump.ReadTypes(),
		"https://cdn.contentful.com",
		os.Getenv("SPACE_ID"),
		os.Getenv("CONTENTFUL_KEY"),
		"en-US",
		impl.FileStore{},
	}

	dumper.Work()

}
