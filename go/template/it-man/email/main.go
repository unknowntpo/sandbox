package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

const text = `Hello {{.UserName}}!
{{range .Emails }}
email {{ . }}
{{- end }}
---------------------
Friends:
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`

func main() {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	// allocate a new template named "test"
	t := template.New("test")

	t = template.Must(t.Parse(text))
	p := Person{UserName: "Eric",
		Emails:  []string{"e850506@gmail.com", "k88036415@gmail.com"},
		Friends: []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
