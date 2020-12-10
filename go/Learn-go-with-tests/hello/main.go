package main

import "fmt"

const english = "English"
const spanish = "Spanish"
const french = "French"
const engHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
