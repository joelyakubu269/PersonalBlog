package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func renderPage(w http.ResponseWriter, filename string, data interface{}) {
	templ,err:= template.ParseFiles("templates/" +filename )
	if err!= nil {
		http.Error(w,"template not found" + err.Error(),http.StatusInternalServerError)
		return
	}
	err= templ.Execute(w,data)
	if err!= nil {
		http.Error(w,"failed to render" + err.Error(),http.StatusInternalServerError)
		return
	}
}
func homeHander(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	articlesDir:= "articles"
	if err:= os.Mkdir(articlesDir,0755); err!= nil { // makes the directory if there are no files in it cause version control deletes empty files
		http.Error(w,"error finding directory" + err.Error(),http.StatusInternalServerError)
		return
	}
	entries,err:= os.ReadDir(articlesDir)
	if err!= nil {
		http.Error(w,"error reading director" + err.Error(),http.StatusInternalServerError)
		return
	}
	var articles []ArticleData
	for _,entry:= range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" { // Used to ignore other directories or non json files
			continue
		}
		filepath:= filepath.Join(articlesDir,entry.Name())
		content, err:= os.ReadFile(filepath)
		if err!= nil {
			continue // skip unreadable files
		}
		err= json.Unmarshal(content,&articles)
		if err!= nil {
			continue // skip corrupted json files 

		}

	}
	renderPage(w,"home.html",articles)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method!= "POST" {
		http.Error(w,"method not allowed",http.StatusInternalServerError)
		return
	}
	id:= r.FormValue("ID") // take note of we assign id and avoid repition and in aspects of deletion
	newId,err := strconv.Atoi(id)
	if err!= nil {
		http.Error(w,"error converting id",http.StatusBadRequest)
		return
	}
	title:= r.FormValue("Title")
	author:= r.FormValue("Author")
	summary:= r.FormValue("Summary")
	content:= r.FormValue("Content")
	val:= ArticleData{
		ID: newId,
		Title: title,
		Author: author,
		Summary: summary,
		Content: content,

	}
	err = saveArticle(val)
	if err!= nil {
		http.Error(w,"unable to create article",http.StatusInternalServerError)
		return
	}
}
func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w,"Status not allowed",http.StatusInternalServerError)
		return
	}
	renderPage(w,"admin.html",nil)
}
func delete(w http.ResponseWriter, r *http.Request){

}