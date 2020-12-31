package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, err := NewLeague(database)
	if err != nil {
		// TODO: handle error
	}
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {

	player := f.league.Find(name)
	if player != nil {
		return player.Wins, nil
	}

	// player not found
	return 0, ErrPlayerNotFound
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		// player not found
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	// Write league back to our database
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
