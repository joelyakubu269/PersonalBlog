package main

import (
	"html/template"
	"net/http"
	"os"
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
	if err:= os.Mkdir(articlesDir,0755); err!= nil {
		http.Error(w,"error finding directory" + err.Error(),http.StatusInternalServerError)
		return
	}
	entries,err:= os.ReadDir(articlesDir)
	if err!= nil {
		http.Error(w,"error reading director" + err.Error(),http.StatusInternalServerError)
		return
	}
	for _,entry:= range
}