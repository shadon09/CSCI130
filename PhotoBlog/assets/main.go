package main

import (
	"html/template"
	"net/http"
	"log"
	"strings"
	"io"
)

var tpl *template.Template

func init(){
	tpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/imgs/", fs)
	http.ListenAndServe(":8080",nil)
}

func index(res http.ResponseWriter, req *http.Request){
	cookie := genCookie(res, req)
	if req.Method == "POST" {
		reader, err := req.MultipartReader()
		if err != nil {
			log.Println("error uploading photo: ", err)
		}
		for{
			part, err := reader.NextPart()
			if err == io.EOF{
				break;
			}
			cookie = uploadPhoto(part, cookie)
			http.SetCookie(res, cookie)
		}
	}
	m := Model(cookie)
	tpl.ExecuteTemplate(res, "index.html", m)
}

func login(res http.ResponseWriter, req *http.Request){
	cookie := genCookie(res, req)
	if req.Method == "POST" && req.FormValue("password") == "secret"{
		m := Model(cookie)
		m.State = true
		m.Name = req.FormValue("name")

		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(m, id)
		http.SetCookie(res, cookie)

		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func logout(res http.ResponseWriter, req *http.Request){
	cookie := newVisitor()
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
}

func genCookie(res http.ResponseWriter, req *http.Request) *http.Cookie{
	cookie, err := req.Cookie("session-fino")
	if err != nil{
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}
	if strings.Count(cookie.Value, "|") != 2{
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}
	if tampered(cookie.Value){
		cookie = newVisitor()
		http.SetCookie(res, cookie)
	}
	return cookie
}
