package main
import(
	"net/http"
)
func main()  {
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/admin",admin)
	http.HandleFunc("/admin")
}