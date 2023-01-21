#!/bin/bash

set -e
set -u

mkdir -p myproject/mymodule/mypackage
cat >myproject/mymodule/mypackage/mypackage.go <<EOF
package mypackage

import "net"

func Get_ips() ([]net.IP, error) {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		return nil, err
	}
	return ips, nil
}
EOF

cat >myproject/mymodule/main.go <<'EOF'
package main

import (
	"html/template"
	"os"

        "mymodule/mypackage"
)

func main() {
	ips, err := mypackage.Get_ips()

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
		IPs:   iplist_str,
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
EOF

pushd myproject/mymodule/
go mod init mymodule
popd

pushd myproject/mymodule/
go build .
popd

./myproject/mymodule/mymodule
