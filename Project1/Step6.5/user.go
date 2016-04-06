package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func newVisitor(res http.ResponseWriter, req *http.Request) *http.Cookie {
	data := InitModel(req.FormValue("name"), req.FormValue("age"))
	code := getCode(data)
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "|" + data + "|" + code,
			// Secure: true,
			HttpOnly: true,
	}
	return cookie
}

func InitModel(name string, age string) string {
	Model := model{
		Name: name,
		Age:  age,
	}

	bs, err := json.Marshal(Model)
	if err != nil {
		fmt.Printf("error: ", err)
	}
	return string(bs)
}