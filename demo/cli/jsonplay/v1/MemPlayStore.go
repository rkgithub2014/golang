package main

//NewPlayStore returns pointer to Memory PlayeStore
func NewPlayStore() *MemPlayStore {
	return &MemPlayStore{map[string]int{}}
}

// NewStoreInstance returns a new instance of Memory Store
func NewStoreInstance() *MemPlayStore {
	s := MemPlayStore{map[string]int{}}
	return &s
}

//MemPlayStore holds data about players in memory
type MemPlayStore struct {
	store map[string]int
}

// RecordWin keeps record of win
func (s *MemPlayStore) RecordWin(name string) {
	s.store[name]++
}

// GetPlayerScore will retrive player's score
func (s *MemPlayStore) GetPlayerScore(name string) int {
	return s.store[name]
}
