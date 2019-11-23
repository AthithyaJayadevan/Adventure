package main

import (
	"cyoa_own/adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "The JSON file that's going to be parsed")
	port := flag.Int("port", 3000, "The port that will use the cyoa application")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := adventure.Jsonparser(f)

	if err != nil {
		panic(err)
	}

	h := adventure.Newhandler(story)
	fmt.Printf("Starting server at : %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
