//Matthew Morado
package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func serveitgood(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func image(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "images/surf.jpg")
}

func design(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "css/design.css")
}

func main() {
	http.HandleFunc("/", serveitgood)
	http.HandleFunc("/image/", image)
	http.HandleFunc("/design/", design)
	http.ListenAndServe(":8080", nil)
}
