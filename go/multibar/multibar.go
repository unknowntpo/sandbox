package main

import (
	"fmt"
	"strings"
	"time"

	"./terminal"
)

const (
	DUMMY_TOTAL_BYTE int = 100
)

type Multibar struct {
	bars []ProgressBar
}

// NewMultibar returns an instance of multibar, with number of bars and length of each bar.
func NewMultibar(numBars int, length int) *Multibar {
	m := &Multibar{
		bars: make([]ProgressBar, numBars),
	}
	for i, _ := range m.bars {
		m.bars[i].length = length
		m.bars[i].totalByte = DUMMY_TOTAL_BYTE
	}
	return m
}

func (m *Multibar) Render() {
	for dummyCounter := 0; dummyCounter < 50; dummyCounter++ {
		// TODO: Wait for all task complete and clean up terminal.
		for i, _ := range m.bars {
			//clearCurrentLine()
			m.bars[i].Render()
			fmt.Print("\n")
		}
		terminal.MoveCursorBeginning(len(m.bars))
		time.Sleep(100 * time.Millisecond)
	}
	terminal.MoveCursorDown(len(m.bars))
}

type ProgressBar struct {
	length      int
	totalByte   int
	currentByte int
}

func (p *ProgressBar) done() bool {
	return p.totalByte <= p.currentByte
}
func (p *ProgressBar) Render() {
	// [====--------]
	if !p.done() {
		fmt.Printf("[")
		fmt.Printf("%s", strings.Repeat("=", p.length*p.currentByte/p.totalByte))
		fmt.Printf("%s", strings.Repeat("-", p.length-p.length*p.currentByte/p.totalByte))
		fmt.Printf("] (%d / %d Bytes)", p.currentByte, p.totalByte)
		// TODO: Decide increment of byte
		p.currentByte += 10
	} else {
		fmt.Printf("[")
		fmt.Printf("%s", strings.Repeat("=", p.length))
		fmt.Printf("] (%d / %d Bytes)", p.currentByte, p.totalByte)
	}
}

func main() {
	multiBar := NewMultibar(5, 50)
	multiBar.Render()
}
