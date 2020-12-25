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

func (i *InMemoryPlayerStore) GetLeague() league {
	var l league
	for name, wins := range i.store {
		l = append(l, Player{Name: name, Wins: wins})
	}
	return l
}
