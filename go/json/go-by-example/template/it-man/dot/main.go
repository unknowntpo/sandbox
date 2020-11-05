package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Eric", 24}
	tmpl, _ := template.New("test").Parse("{{.}}, Name: {{.Name}}, Age: {{.Age}}\n")
	_ = tmpl.Execute(os.Stdout, p)
}
