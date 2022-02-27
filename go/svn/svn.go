package main

import "fmt"

type users map[userID]user

type userID int64

type user struct {
	id    userID // userid
	name  string // name of sheet
	email string
	sheet *sheet
}

type cell string

type sheet struct {
	cache   [][]cell // cache stores the newest data for sheet.
	commits []commit
}

func newSheet(table [][]row, commits []commit) {
	sheet := &sheet{
		cache: make([][]cell, 0, 0),
	}

	// copy values from table to cache
	for _, rows := range table {
		copy(sheet.cache, rows)
	}
	return sheet
}

type commit struct {
	id         string // hash or uuid
	diffs      []diff
	created_at time.Time // timestamp
}

type diff struct {
	position
	minus
	plus
}

func main() {
	fmt.Println("vim-go")
}
