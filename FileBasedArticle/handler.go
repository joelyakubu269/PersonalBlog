package main
import(
	"net/http"
	"html/template"

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
}