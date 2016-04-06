package main

import(
	"encoding/json"
	"fmt"
	"strings"
)

type model struct {
	Name string
	Age  string
}

func Model(s string) model{
	xs := strings.Split(s, "|")
	usrData := xs[1]
	var m model
	err := json.Unmarshal([]byte(usrData), &m)
	if err != nil{
		fmt.Printf("error unmarshalling: %v", err)
	}
	return m;
}