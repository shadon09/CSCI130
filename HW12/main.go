package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("my-session")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "my-session",
				Value: id.String(),
				// Secure: true,
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)
		}
		fmt.Printf("%v", cookie)
	})
	http.ListenAndServe(":8080", nil)
}
