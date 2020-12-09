package main

import (
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

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
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
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Make our router to handle routing
	p.router.ServeHTTP(w, r)
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// Create a Server with store and router
	p := &PlayerServer{
		store:  store,
		router: http.NewServeMux(),
	}
	// register our handler into our mux manually
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
