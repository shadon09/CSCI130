package main

import(
	"io"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/", func(res http.ResponsWriter, req *http.Request){
		if req.URL.Path != "/"{
			http.NotFound(res, req)
			return
		}
		cookie, err:= req.Cookie("my-cookie")
		if err == http.ErrNo
	}
}
