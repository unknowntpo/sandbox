package main

type Greeter struct {
	name string
}

func (g Greeter) Info() string {
	return "This is a greeter !"
}

func (g Greeter) Register() Module {
	return &Greeter{name: "greeter"}
}

func main() {

}
