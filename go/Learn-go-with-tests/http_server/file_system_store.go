package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	// TODO: Handle dummy return values
	league := f.GetLeague()
	for _, player := range league {
		if player.Name == name {
			return player.Wins, nil
		}
	}
	// player not found
	return 0, ErrPlayerNotFound

}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	// TODO: record wins toleauge table
	league := f.GetLeague()

	hasPlayer := false
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
			//player.Wins++
			hasPlayer = true
		}
	}

	// player not found
	// same underlying array?
	if hasPlayer == false {
		league = append(league, Player{Name: name, Wins: 1})
	}

	// Write league back to our database
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetLeague() league {
	f.database.Seek(0, 0)
	// TODO: Handle errs
	league, _ := NewLeague(f.database)
	return league
}
