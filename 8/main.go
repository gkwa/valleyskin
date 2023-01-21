package main

import (
	"fmt"
	"net"
	"os"
	"text/template"
)

func main() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("google.com. IN A %s\n", ip.String())
	}
}

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{
		Name:        "Test templates",
		Description: "Let's test a template to see the magic.",
	}

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
