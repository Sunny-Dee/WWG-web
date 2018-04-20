package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/jtacoma/uritemplates"
)

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

	// this can error due to reasons that are hard
	// to produce in tests such as network resets
	// on the given socket
	// respBody, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Println("response body", response.Body)
	c := &Comic{}
	json.NewDecoder(response.Body).Decode(c)

	fmt.Println("comic title:", c.Title)

	return c, nil

}
