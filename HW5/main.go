package main

import (
	"html/template"
	"log"
	"net/http"
)

func quote(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", quote)
	http.ListenAndServe(":8080", nil)
}
