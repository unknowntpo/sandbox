package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {

	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins, nil
	}

	// player not found
	return 0, ErrPlayerNotFound
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		// player not found
		league = append(league, Player{Name: name, Wins: 1})
	}

	// Write league back to our database
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	// TODO: Handle errs
	league, _ := NewLeague(f.database)
	return league
}
