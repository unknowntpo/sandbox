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

// Find finds player by name in league
// Return pointer to player if player is found.
// If player not found, return nil
func (l league) Find(name string) *Player {

	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}
