package main

import (
	"./dump"
	"os"
)

func main() {

	dumper := dump.Dumper{
		dump.ReadTypes(),
		"dump.ReadTypes()",
		os.Getenv("SPACE_ID"),
		os.Getenv("CONTENTFUL_KEY"),
	}


	dumper.Work()

}
