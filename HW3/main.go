package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name  string
	Major string
}

type student struct {
	person
	Graduate bool
}

func main() {
	p1 := student{
		person: person{
			Name:  "Matthew Morado",
			Major: "Computer Science",
		},
		Graduate: false,
	}
	tpl, err := template.ParseFiles("conditional.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
