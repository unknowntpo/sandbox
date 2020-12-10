package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func calculate(e []string) (float64, error) {
	result := 0.0
	num1, err := strconv.ParseFloat(e[0], 64)
	if err != nil {
		return 0.0, err
	}
	num2, err := strconv.ParseFloat(e[2], 64)
	if err != nil {
		return 0.0, err
	}
	switch e[1] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero.")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	default:
		return 0.0, errors.New("error: don't know how to do that")
	}
	return result, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("error: wrong number of arguments")
		return
	}
	res, err := calculate(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
