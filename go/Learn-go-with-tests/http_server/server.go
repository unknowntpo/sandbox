package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrPlayerNotFound = errors.New("Player not found")
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string // append this slice by winner's name when someone wins
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
}

// Player stores a name with a number of wins.
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store        PlayerStore
	http.Handler // Structure embedding
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.scores[name]
	if !ok {
		return 0, ErrPlayerNotFound
	}
	return score, nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	// Create a Server with store and router
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, err := p.store.GetPlayerScore(player)

	// Handle the not founded player
	if err == ErrPlayerNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Player not found")
		return
	}
	fmt.Fprint(w, score)
}
