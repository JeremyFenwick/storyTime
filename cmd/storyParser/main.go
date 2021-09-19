package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JeremyFenwick/storyTime"
)

func main() {
	port := flag.Int("port", 3000, "The port number to be used: ")
	flag.Parse()

	storyFlag := flag.String("story", "story.json", "Enter json file with the story:")
	flag.Parse()
	story := *storyFlag

	fmt.Printf("Using story %s...\n", story)

	file, err := os.Open(story)
	if err != nil {
		fmt.Println(err)
	}

	narrative, err := storyTime.JsonStory(file)
	if err != nil {
		fmt.Println(err)
	}

	h := storyTime.NewHandler(narrative)
	fmt.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
