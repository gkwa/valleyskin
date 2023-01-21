package main

import (
	"html/template"
	"os"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tpl1 := template.ParseFiles("layout.html")
	tmpl := template.Must(tpl1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}

func main2() {
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
