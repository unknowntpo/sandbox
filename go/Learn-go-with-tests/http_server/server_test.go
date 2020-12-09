package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

// Test the functionality of InMemoryPlayerStore by sending POST /players/Pepper request
// for 3 times, and use GET /players/Pepper request to check the statuscode,
// and result of score recording.
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{
		map[string]int{},
	}
	server := NewPlayerServer(&store)

	player := "Pepper"

	// Send POST /players/Pepper request for 3 times, expect score of Pepper is equal to 3
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	// Send GET /players/Pepper request to check the score recording result
	req := newGetScoreRequest(player)
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	assertStatus(t, resp.Code, http.StatusOK)
	assertResponseBody(t, resp.Body.String(), "3")
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)
	t.Run("It returns 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, req)

		assertStatus(t, resp.Code, http.StatusOK)
	})
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
