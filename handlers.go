package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
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
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, xkcd)
}
