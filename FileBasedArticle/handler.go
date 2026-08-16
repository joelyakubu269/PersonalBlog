package main

import (
	"encoding/json"
	"fmt"

	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func renderPage(w http.ResponseWriter, filename string, data interface{}) {
	templ, err := template.ParseFiles("templates/" + filename)
	if err != nil {
		http.Error(w, "template not found"+err.Error(), http.StatusInternalServerError)
		return
	}
	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to render"+err.Error(), http.StatusInternalServerError)
		return
	}
}
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("there is an err in article")
		return
	}
	idstr := r.URL.Query().Get("id")

	article, err := readArticle(idstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderPage(w, "article.html", article)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("there is an err in home")
		return
	}
	articlesDir := "articles"
	if err := os.MkdirAll(articlesDir, 0755); err != nil { // makes the directory if there are no files in it cause version control deletes empty files
		http.Error(w, "error finding directory"+err.Error(), http.StatusInternalServerError)
		return
	}
	entries, err := os.ReadDir(articlesDir)
	if err != nil {
		http.Error(w, "error reading director"+err.Error(), http.StatusInternalServerError)
		return
	}
	var articles []ArticleData
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" { // Used to ignore other directories or non json files
			continue
		}
		filepath := filepath.Join(articlesDir, entry.Name())
		content, err := os.ReadFile(filepath)
		var article ArticleData
		if err != nil {
			continue // skip unreadable files
		}
		err = json.Unmarshal(content, &article)
		if err != nil {
			continue // skip corrupted json files

		}
		articles = append(articles, article)

	}
	renderPage(w, "home.html", articles)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Show the form
		renderPage(w, "createArticle.html", nil)

	case http.MethodPost:

		id, err := generateNextArticleID()
		if err != nil {
			http.Error(w, "error generating id", http.StatusInternalServerError)
			return
		}

		title := r.FormValue("Title")
		author := r.FormValue("Author")
		summary := r.FormValue("Summary")
		content := r.FormValue("Content")
		val := ArticleData{
			ID:      id,
			Title:   title,
			Author:  author,
			Summary: summary,
			Content: content,
		}
		err = saveArticle(val)
		if err != nil {
			http.Error(w, "unable to create article", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Status not allowed", http.StatusMethodNotAllowed)
		fmt.Println("there is an err in admin")
		return
	}
	articles, err := getArticles()
	if err != nil {
		http.Error(w, "unable to get articles", http.StatusInternalServerError)
		return
	}
	renderPage(w, "admin.html", articles)
}
func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("there is an err in delete")
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("ID:", id)
	val, err := idConverter(id)
	if err != nil {
		http.Error(w, "status bad request", http.StatusBadRequest)
		return
	}
	err = deleteArticle(val)
	if err != nil {
		http.Error(w, "unable to delete file", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
func editHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Show the form
		        id := r.URL.Query().Get("id")
				 fmt.Println("ID:", id)


        article, err := readArticle(id)
        if err != nil {
            http.Error(w, "unable to find article", http.StatusNotFound)
            return
        }

		renderPage(w, "edit.html", article)
		fmt.Println(err)

	case http.MethodPost:
		
	id := r.URL.Query().Get("id")
	idn, err := idConverter(id)
	if err != nil {
		http.Error(w, "status bad request", http.StatusBadRequest)
		return
	}
	title := r.FormValue("Title")
	author := r.FormValue("Author")
	summary := r.FormValue("Summary")
	content := r.FormValue("Content")
	val := ArticleData{
		ID:      idn,
		Title:   title,
		Author:  author,
		Summary: summary,
		Content: content,
	}
	err = saveArticle(val)
	if err != nil {
		http.Error(w, "unable to create article", http.StatusInternalServerError)
		return
	}
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
	return
}
}
