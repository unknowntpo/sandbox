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
	store PlayerStore
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
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	// Dummy name, modify it later
	p.store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, err := p.store.GetPlayerScore(player)

	// Handle the not founded player
	if err == ErrPlayerNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Player not found")
		return
	}
	fmt.Fprint(w, score)
}
