package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	// data for our test database
	initData := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// call store.GetLeague() again to check seek problem
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got, _ := store.GetPlayerScore("Chris")

		want := 34
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for non-existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")

		want := 1
		assertScoreEquals(t, got, want)
	})
	t.Run("get player not found error", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		_, err := store.GetPlayerScore("Pepper")

		if err == nil {
			t.Errorf("expect an error, but got nothing")
		}

		if err != ErrPlayerNotFound {
			t.Errorf("wrong error message: got %v, want %v", ErrPlayerNotFound, err)
		}
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
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
