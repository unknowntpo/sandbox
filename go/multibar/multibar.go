package main

import (
	"fmt"
	"strings"
	"time"
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
	for {
		// TODO: jump back to beginning of bar 0
		for i, _ := range m.bars {
			m.bars[i].Render()
		}
		time.Sleep(100 * time.Millisecond)
	}
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
	multiBar := NewMultibar(5, 10)
	multiBar.Render()
}
