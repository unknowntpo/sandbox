package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	jsonContentType = "application/json"
)

var (
	ErrPlayerNotFound = errors.New("Player not found")
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string // append this slice by winner's name when someone wins
	league   league   // Store some players in our store
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
	GetLeague() league
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

func (s *StubPlayerStore) GetLeague() league {
	return s.league
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
	// Why we have to set header of response before encode json?
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
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
