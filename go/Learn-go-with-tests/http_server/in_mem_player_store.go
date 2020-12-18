package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	// How to init a map?
	return &InMemoryPlayerStore{store: map[string]int{}}
}
func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return i.store[name], nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
