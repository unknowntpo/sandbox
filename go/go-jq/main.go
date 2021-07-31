package main

import "fmt"

const in = `{"task": {"id": 1,"title": "hello","content": "do some homework","done": true,"version": 1}}`

const defaultBufSize = 100

type Formatter struct {
	indent int
	out    []byte // the buffer where we store the modified output.
}

func NewFormatter(indent int) *Formatter {
	return &Formatter{indent: indent, out: []byte{}}
}

func (f *Formatter) Format(in []byte) {
	depth := 0

	for _, v := range in {
		switch v {
		case ',':
			f.out = append(f.out, []byte(fmt.Sprintf(",\n%*s", depth*f.indent, ""))...)
		case '{':
			depth++
			f.out = append(f.out, []byte(fmt.Sprintf("{\n%*s", depth*f.indent, ""))...)
		case '}':
			depth--
			f.out = append(f.out, []byte(fmt.Sprintf("\n%*s}", depth*f.indent, ""))...)
		case '[':
			f.out = append(f.out, []byte(fmt.Sprintf("\n%*s[", depth*f.indent, ""))...)
			depth++
		case ']':
			depth--
			f.out = append(f.out, []byte(fmt.Sprintf("\n%*s]", depth*f.indent, ""))...)
		default:
			f.out = append(f.out, v)
		}

	}
}
func main() {
	f := NewFormatter(4)
	f.Format([]byte(in))
	fmt.Println(string(f.out))
}
