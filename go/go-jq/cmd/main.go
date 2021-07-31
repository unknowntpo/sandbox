package main

import (
	"bufio"
	"fmt"
	"os"
)

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
			depth++
			f.out = append(f.out, []byte(fmt.Sprintf("[\n%*s", depth*f.indent, ""))...)
		case ']':
			depth--
			f.out = append(f.out, []byte(fmt.Sprintf("\n%*s]", depth*f.indent, ""))...)
		default:
			f.out = append(f.out, v)
		}

	}
}

// Clear clears the output buffer.
func (f *Formatter) Clear() {
	f.out = []byte{}
}

func main() {
	/*
		inputs := []string{
			`{"task": {"id": 1,"title": "hello","content": "do some homework","done": true,"version": 1}}`,
			`{"messages": ["hello","hello","hello"]}`,
		}

		f := NewFormatter(4)
		for _, in := range inputs {
			f.Format([]byte(in))
			fmt.Println(string(f.out))
			f.Clear()
			fmt.Println()
		}
	*/
	var reader = bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')

	f := NewFormatter(4)
	f.Format([]byte(in))
	fmt.Println(string(f.out))
	f.Clear()
	fmt.Println()
}
