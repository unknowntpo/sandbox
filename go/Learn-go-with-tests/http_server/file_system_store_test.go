package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		initData := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// call store.getleague() again to check seek problem
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

// createTempFile create a temporary file with initial data specified by caller,
// and return that file with the clean up function removeFile, caller can use it to
// clean up temporary file easilly with defer keyword.
func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)

	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
