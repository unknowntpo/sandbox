package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/sethgrid/curse"
)

func ProgressBar() {
	fmt.Println("Progress Bar")
	total := 150
	progressBarWidth := 80
	c, _ := curse.New()

	NumBars := 3
	for i := 0; i < NumBars; i++ {
		// give some buffer space on the terminal
		fmt.Println()

		// display a progress bar
		for i := 0; i <= total; i++ {
			c.MoveUp(1)
			c.EraseCurrentLine()
			fmt.Printf("%d/%d ", i, total)

			c.MoveDown(1)
			c.EraseCurrentLine()
			fmt.Printf("%s", progressBar(i, total, progressBarWidth))

			time.Sleep(time.Millisecond * 25)
		}
		// end the previous last line of output
		fmt.Println()
	}
	fmt.Println("Complete")
}

func progressBar(progress, total, width int) string {
	bar := make([]string, width)
	for i, _ := range bar {
		// decide the content in that in bar[i]
		if float32(progress)/float32(total) > float32(i)/float32(width) {
			bar[i] = "*"
		} else {
			bar[i] = " "
		}
	}
	return "[" + strings.Join(bar, "") + "]"
}

func main() {
	ProgressBar()
}
