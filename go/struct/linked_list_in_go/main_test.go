package main

import (
     "testing"
)

func TestCreatePlayList(t *testing.T) {
	listName := "testing"

	testList := createPlayList(listName)
	if testList.head != nil {
		t.Errorf("Fail on create new playlist")
	}
}
