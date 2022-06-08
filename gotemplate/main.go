package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/hello.html"))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "hello.html", nil)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/css/", fs)
	// http.Handle("/js/", fs)
	http.ListenAndServe(":9999", nil)
}
