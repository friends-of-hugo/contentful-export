package main

import "./dump"

func main() {

	types := dump.ReadTypes()
	dump.Work(types)

}
