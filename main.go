package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	port    string
	baseURL string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load environment variables
	port = os.Getenv("PORT")
	baseURL = os.Getenv("XKCD_BASE_URL")
}

func main() {
	// Setup your handlers!
	http.HandleFunc("/", index)
	http.HandleFunc("/rendercomic", renderComic)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Start up your server!
	fmt.Printf("Starting program.\nListening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
