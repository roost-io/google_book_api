package main

import (
	"embed"
	"flag"
	"html/template"
	"net/http"
)

//go:embed templates
var res embed.FS
var tmpl *template.Template
var apiEndpoint string

func init() {
	tmpl = template.Must(template.ParseFS(res, "templates/hello.html"))
	flag.StringVar(&apiEndpoint, "apiEndpoint", "localhost:8081", "a string")
	flag.Parse()
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	variableMap := map[string]interface{}{
		"apiEndpoint": apiEndpoint,
	}
	tmpl.ExecuteTemplate(w, "hello.html", variableMap)
}

func main() {

	http.FileServer(http.FS(res))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8084", nil)
}
