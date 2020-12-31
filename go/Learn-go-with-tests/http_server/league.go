package main

import (
	"encoding/json"
	"io"
)

type League []Player

// NewLeague creates a League from rdr, which stores JSON data.
func NewLeague(rdr io.Reader) (League, error) {
	var l League
	err := json.NewDecoder(rdr).Decode(&l)
	return l, err
}

// Find finds player by name in League
// Return pointer to player if player is found.
// If player not found, return nil
func (l League) Find(name string) *Player {

	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}
