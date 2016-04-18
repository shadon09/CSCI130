package main

import (
	"html/template"
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"os"
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
		src, hdr, err := req.FormFile("data")
		if err != nil {
			log.Println("error uploading photo: ", err)
		}
		cookie = uploadPhoto(src, hdr, cookie)
		http.SetCookie(res, cookie)
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
		exists, _ := userExists("assets/imgs/" + m.Name)
		if exists == false{
			os.MkdirAll("assets/imgs/" + m.Name, os.ModePerm)
		}
		xs := strings.Split(cookie.Value, "|")
		id := xs[0]

		cookie := currentVisitor(m, id)
		cookie = getPhotos(cookie)
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
func getPhotos(c *http.Cookie) *http.Cookie{
	cookie := c
	m := Model(cookie)
	files, _ := ioutil.ReadDir("./assets/imgs/" + m.Name + "/")
	for _, f := range files {
		m.Pictures = append(m.Pictures, "/imgs/"+ m.Name + "/" + f.Name())
		xs := strings.Split(c.Value, "|")
		id := xs[0]
		cookie = currentVisitor(m, id)
	}
	return cookie
}
