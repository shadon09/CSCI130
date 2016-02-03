package main

import (
	"io"
	"net/http"
)

func quote(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "'Placeholder Quote'\n")
	io.WriteString(res, " -Me\n")
	io.WriteString(res, "Matthew Morado")
}

func main() {
	http.HandleFunc("/", quote)
	http.ListenAndServe(":8080", nil)
}
