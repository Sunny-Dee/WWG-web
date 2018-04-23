package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtacoma/uritemplates"
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func renderComic(w http.ResponseWriter, r *http.Request) {
	xkcd, err := fetchComic()
	if err != nil {
		xkcd = &Comic{
			Title:  "Not a real title",
			Year:   "1969",
			ImgURL: "not a real url",
		}
	}
	fmt.Println(xkcd.Title)
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, xkcd)
}

const latestComicNum = 1000

type Comic struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImgURL string `json:"img"`
}

func fetchComic() (*Comic, error) {
	comicNum := rand.Intn(latestComicNum)
	fmt.Println("Generated random comic number to retrieve:", comicNum)
	template, err := uritemplates.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	url, err := template.Expand(map[string]interface{}{"comicNo": comicNum})
	if err != nil {
		return nil, err
	}

	fmt.Println("get request url", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status response: %d", response.StatusCode)
	}

	fmt.Println("response body", response.Body)
	c := &Comic{}
	json.NewDecoder(response.Body).Decode(c)

	fmt.Println("comic title:", c.Title)

	return c, nil

}
