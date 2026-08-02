package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	idstr := r.URL.Query().Get("id")
	Id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Fprint(w, "bad request", http.StatusBadRequest)
		return
	}
	templ, err := template.ParseFiles("article.html")

	for _, article := range articles {
		if article.ID == Id {
		}
	}
}
