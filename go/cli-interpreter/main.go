package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseArgs() {
	// wait for user input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in := scanner.Text()
		inputSlice := strings.Split(in, " ")
		mux(inputSlice)
	}
}
func showPrompt() {
	fmt.Printf("cmd>")
}

func cli_Init() {
	for {
		showPrompt()
		parseArgs()
	}
}
func main() {
	cli_Init()
}
