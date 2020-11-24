package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// arg of NewReq?
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// What is httptest ? newrecorder?
		resp := httptest.NewRecorder()

		// Call the handlefunc to get response for testing
		PlayerServer(resp, req)

		// What does string() mean?
		got := resp.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
