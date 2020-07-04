package main

//InMemoryStore to keep track of player score
type InMemoryStore struct {
	store map[string]int
}

// NewInMemoryStore return instance of store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}

//GetPlayerScore returns players score
func (m *InMemoryStore) GetPlayerScore(player string) int {
	return m.store[player]
}

//RecordWin records players win
func (m *InMemoryStore) RecordWin(player string) {
	m.store[player]++
}

//ShowScoreAll returns a map of all players
func (m *InMemoryStore) ShowScoreAll() map[string]int {
	return m.store
}
