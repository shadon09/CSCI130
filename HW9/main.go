package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "n"
		val := req.FormValue(key)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
		<input type="text" name="n">
		<input type="submit">
		</form>`+val)
	})
	http.ListenAndServe(":8080", nil)
}
