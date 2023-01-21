package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	u := User{"John", "a regular user"}

	ut, err := template.New("users").Parse("The user is {{ .Name }} and he is {{ .Bio }}.")
	if err != nil {
		panic(err)
	}

	err = ut.Execute(os.Stdout, u)

	if err != nil {
		panic(err)
	}
}
