package main

import (
	"fmt"
	"html/template"
	"net"
	"os"
)

func main() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}

	multiline := `
	<p>My stuff header {{ .Name }}</p>
	<h2>My loop over IPs:</h2>
	{{ range .IPs -}}
	<p>{{ . }}</p>
	{{ end }}
	`

	var iplist_str []string

	for _, ip := range ips {
		iplist_str = append(iplist_str, ip.String())
	}

	data := struct {
		Name string
		IPs  []string
	}{
		Name: "John Doe",
		IP:   iplist_str,
	}

	tpl, err := template.New("stuff").Parse(multiline)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
