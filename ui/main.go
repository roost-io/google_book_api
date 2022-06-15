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
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/css/", http.StripPrefix("/css", fs))

	// http.Handle("/", http.FileServer(http.Dir("css/")))
	// fs := http.FileServer(http.Dir("templates"))
	// http.Handle("/css/", fs)
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates"))))
	// http.Handle("/", http.FileServer(http.Dir("css/")))
	// http.Handle("/ui/", http.StripPrefix("templates/css/", http.FileServer(http.Dir("templates/css"))))
	// http.HandleFunc("/css/", func(response http.ResponseWriter, request *http.Request) {
	// 	http.ServeFile(response, request, "/templates/css/")
	// })

	http.ListenAndServe(":8084", nil)
}
