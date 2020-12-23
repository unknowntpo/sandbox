package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	// TODO: Handle errs
	league, _ := NewLeague(f.database)
	return league
}
