package main

import (
	"encoding/json"
	"io"
)

type league []Player

// NewLeague creates a league from JSON.
func NewLeague(rdr io.Reader) (league, error) {
	var league league
	err := json.NewDecoder(rdr).Decode(&league)
	return league, err
}
