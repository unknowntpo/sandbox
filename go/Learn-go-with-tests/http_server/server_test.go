package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
			"MMD":    0,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)

		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)

		assertResponseBody(t, resp.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusNotFound)
	})

	t.Run("returns statusCode 200 and score 0 on existing players with score 0", func(t *testing.T) {
		req := newGetScoreRequest("MMD")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "0")
	})

}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it wins when POST", func(t *testing.T) {
		player := "Pepper"
		req := newPostWinRequest(player)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		// Check if name recorded is correct
		if store.winCalls[len(store.winCalls)-1] != player {
			t.Errorf("did not store correct winner, got %q, want %q", store.winCalls[len(store.winCalls)-1], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("It returns the league table as JSON", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)
		assertStatus(t, resp.Code, http.StatusOK)
		assertContentType(t, resp, jsonContentType)
		assertLeague(t, got, wantedLeague)
	})
}

func BenchmarkLeague(b *testing.B) {
	b.Run("get league from InMemoryPlayerStore", func(b *testing.B) {
		store := InMemoryPlayerStore{map[string]int{
			"Cleo":  32,
			"Chris": 20,
			"Tiest": 14,
		}}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = store.GetLeague()
		}
	})

	b.Run("get league from StubPlayerStore", func(b *testing.B) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = store.GetLeague()
		}
	})
	b.Run("get league from FileSystemPlayerStore", func(b *testing.B) {
		initData := `[
            {"Name": "Cleo", "Wins": 32},
            {"Name": "Chris", "Wins": 20},
	    {"Name": "Tiest", "Wins": 14}]`

		database, cleanDatabase := createTempFileForBench(b, initData)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = store.GetLeague()
		}
	})

}

func createTempFileForBench(tb testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	tb.Helper()
	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		tb.Fatalf("could not create temp file %v", err)

	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

// Some helper functions:
func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// Return the request for given name
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Unexpected status, got %d, want %d", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
func assertLeague(t *testing.T, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, want string) {
	t.Helper()
	if resp.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, resp.Result().Header)
	}
}
func getLeagueFromResponse(t *testing.T, body io.Reader) (l League) {
	t.Helper()

	l, err := NewLeague(body)
	if err != nil {
		t.Fatalf("Unable to parse response %q from server into slice of Player, '%v'", body, err)
	}

	return l

}
