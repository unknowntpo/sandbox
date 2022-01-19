package main

import "fmt"

type Module interface {
	Register() Module
	Info() string
	Do() error
}

type modules map[string]Module
