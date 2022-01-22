package main

import (
	"fmt"
)

type Add struct{}

func (a *Add) Info() string { return "calculator.add" }

type Sub struct{}

func (s *Sub) Info() string { return "calculator.sub" }

var modules = make(map[string]Module)

func RegisterModule(m Module) {
	name := m.Info()
	if _, exist := modules[name]; exist {
		fmt.Println("ERROR: Module %s already exist", name)
		return
	}
	modules[name] = m
}

type Module interface {
	Info() string
}

func init() {
	RegisterModule(&Add{})
	RegisterModule(&Sub{})
}

func main() {
	fmt.Println(modules)
}
