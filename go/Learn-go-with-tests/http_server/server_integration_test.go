package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
