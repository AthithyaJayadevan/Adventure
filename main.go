package main

import (
	"cyoa_own/adventure"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "The JSON file that's going to be parsed")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := adventure.Jsonparser(f)

	if err != nil {
		panic(err)
	}

	//REFACTORING

	fmt.Printf("%+v\n", story)

}
