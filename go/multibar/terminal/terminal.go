package terminal

import (
	"fmt"
)

const (
	ESC = "\033"
	CSI = ESC + "["
)

// moves cursor to beginning of n lines up.
func MoveCursorBeginning(n int) {
	fmt.Printf("%s%dF", CSI, n)
}

// moves cursor n lines down.
func MoveCursorDown(n int) {
	fmt.Printf("%s%dB", CSI, n)
}

// moves cursor left n columns.
func MoveCursorBack(n int) {
	fmt.Printf("%s%dD", CSI, n)
}

func clearCurrentLine() {
	fmt.Print(CSI + "K")
}

func main() {
	fmt.Printf("01234567890\n")
	fmt.Printf("01234567890\n")
	//MoveCursorBack(2)
	MoveCursorBeginning(2)
	fmt.Print("@")
}
