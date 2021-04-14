package main

import (
	"os"
	"text/template"
)

const text = `{{23 -}} < {{- 45}}`

func ExampleWhiteSpace() {
	t := template.Must(template.New("t1").Parse(text))
	t.Execute(os.Stdout, nil)
	// Output:
	// 23<45

}
