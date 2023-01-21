package main

import (
	"html/template"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name   string
		Skills []string
	}{
		Name: "John Doe",
		Skills: []string{
			"C++",
			"Java",
			"Python",
		},
	}

	err = tpl.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
