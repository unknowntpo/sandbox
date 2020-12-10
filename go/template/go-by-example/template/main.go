package main

import (
	"fmt"
	"os"
	"text/template"
)

var card1 = map[int]string{1: "S", 2: "H", 3: "C", 4: "D"}
var card2 = map[string]string{"a": "S", "b": "H",
	"c": "C", "d": "D"}

const tmpl1 = `Dot:{{.}}`
const tmpl2 = `Hi, {{.a}}`
const tmpl3 = `For {{ range $k, $v := . }}
k = {{printf "%03d" $k}} v = {{$v}}
{{- end}}
`

const tmpl4 = `define{{define "T1" -}}
Apple 
{{- end}}end 
define{{define "T2" -}}
Ape
{{- end}}end
{{template "T2"}} ate {{template "T1"}}
`

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Display template t1")

	t := template.Must(template.New("t1").Parse(tmpl1))
	check(t.Execute(os.Stdout, card1))

	// Why these failed?
	// Key should be alphanumeric
	/*
		t := template.Must(template.New("greet").Parse("Hi,{{.1}}\n"))
		check(t.Execute(os.Stdout, card1))
	*/

	fmt.Println("\nDisplay template t2")
	t = template.Must(template.New("t2").Parse(tmpl2))
	check(t.Execute(os.Stdout, card2))

	fmt.Println("\nDisplay template t3")

	// Use variable to loop through map
	t = template.Must(template.New("t3").Parse(tmpl3))
	check(t.Execute(os.Stdout, card1))

	fmt.Println("----------")
	// Use logic function to display repeated content
	t = template.Must(template.New("t4").Parse(tmpl4))
	check(t.Execute(os.Stdout, card1))

}
