package main

import (
	"net/http"
	"strconv"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	Id, err := strconv.Atoi
	for _, article := range articles {
		if article.ID == Id {
			templ.Execute.ParseFiles("article.html")
		}
	}
}
