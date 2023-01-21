package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	multiline := `You have a task:
   name: "{{ .Name}}"
   description: "{{ .Description}}"
`

	tpl, err := template.New("todos").Parse(multiline)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}
