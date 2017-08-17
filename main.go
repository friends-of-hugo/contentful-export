package main

import (
	"os"

	"./dump"
)

func main() {

	dumper := dump.Dumper{
		dump.ReadTypes(),
		"https://cdn.contentful.com",
		os.Getenv("SPACE_ID"),
		os.Getenv("CONTENTFUL_KEY"),
		"en-US",
	}

	dumper.Work()

}
