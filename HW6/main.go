package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, req.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
