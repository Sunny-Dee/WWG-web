package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/jtacoma/uritemplates"
)

const latestComicNum = 1000

// Comic shapes the response from the XKCD api into a struct that
// we can consume.
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
