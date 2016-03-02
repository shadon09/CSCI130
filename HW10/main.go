package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	var s string
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			src, _, err := req.FormFile("file")
			if err != nil {
				panic(err)
			}
			dst, err := os.Create(filepath.Join("./", "file.txt"))
			defer src.Close()
			io.Copy(dst, src)
			temp, err := ioutil.ReadFile("file.txt")
			s = string(temp)
		}
		err = tpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Println(err)
		}
		io.WriteString(res, `<br>`+s)

	})
	http.ListenAndServe(":8080", nil)
}
