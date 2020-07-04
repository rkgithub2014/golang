package main

import (
	"fmt"
	"net/http"
)

//

//PlayerStore provides metods to access PlayerStore
type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
	ShowScoreAll() map[string]int
}

//PlayerServer HTTP interface for PlayerStore
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

//NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		if player == "all" {
			p.showScoreAll(w)
		} else {
			p.showScore(w, player)
		}

	case http.MethodPost:
		p.processWin(w, player)
	}
}

/* func (p *PlayerServer) ServeHTTP(respw http.ResponseWriter, request *http.Request) {

	player := request.URL.Path[len("/players/"):]

	switch request.Method {
	case http.MethodGet:
		if player == "all" {
			p.showScoreAll(respw)
		} else {
			p.showScore(respw, player)
		}

	case http.MethodPost:
		p.processWin(respw, player)
	}
}
*/
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, "player:%v, Score:%v", player, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	if player != "all" {
		p.store.RecordWin(player)
	}
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScoreAll(w http.ResponseWriter) {
	fmt.Fprintf(w, "player: Score:")
	for key, value := range p.store.ShowScoreAll() {
		fmt.Fprintf(w, "%v,%v", key, value)
	}
}
