package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fs := strings.Split(req.URL.Path, "/")
		s := strings.Join(fs, " - ")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, s)
	})
	http.ListenAndServe(":8080", nil)
}
