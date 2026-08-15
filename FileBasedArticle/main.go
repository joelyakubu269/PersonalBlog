package main

import (
	"fmt"
	"net/http"
)
func main()  {
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/admin",admin)
	http.HandleFunc("/article",ArticleHandler)
	http.HandleFunc("/create",createHandler)
	http.HandleFunc("/admin/article/edit",editHandler)
	http.HandleFunc("/admin/article/delete",delete)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server is up and running")
}