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


package main

import (
	"html/template"
	"text/template"
	"os"
)

func main() {
	multiline := `
<h1>Go templates</h1>
<p>The user is {{ .Name }}</p>
<h2>Skills:</h2>
{{ range .Skills }}
    <p>{{ . }}</p>
{{ end }}
`

	type IPs struct {
		IP string
	}

	tpl, err := template.New("mystuff").Parse(multiline)

	if err != nil {
		panic(err)
	}

	data := struct {
		Name   string
		IPs []string
	}{
		Name: "John Doe",
		IPs: []string{
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
