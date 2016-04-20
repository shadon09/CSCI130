package main

import(
	"html/template"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl, _ = template.ParseGlob("*.html")
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request){
	tpl.ExecuteTemplate(res,"index.html", nil)
}
