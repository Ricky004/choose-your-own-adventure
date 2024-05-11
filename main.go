package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ricky004/choose-your-own-adventure/story"
)

func main() {
	port := flag.Int("port", 3000, "the port to the server")
	file := flag.String("file", "story/story.json", "the JSON file")
	flag.Parse()
	fmt.Printf("Using the story %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	
    s, err := story.JSONStory(f)
	if err != nil {
		panic(err)
	}  
	
	h := story.NewHandler(s)
	fmt.Printf("Starting the server: %d\n",*port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h ))
}
