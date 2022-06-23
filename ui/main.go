package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates
var res embed.FS
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFS(res, "templates/hello.html"))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "hello.html", nil)
}

func main() {
	http.FileServer(http.FS(res))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8084", nil)
}
