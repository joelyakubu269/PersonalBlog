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
	templ, err := template.ParseFiles("templates/article.html")
	if err != nil {
		fmt.Fprint(w, "internal server error", http.StatusInternalServerError)
		return
	}

	for _, article := range articles {
		if article.ID == Id {
			templ.Execute(w, article)
			return
		}
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Fprint(w, "internal server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	templ.Execute(w, articles)

}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/article", ArticleHandler)
	http.ListenAndServe(":8080", nil)
}
