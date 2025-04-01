package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string `json:"Name"`
	Wins int    `json:"Wins"`
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, _ *http.Request) {
	leagueTable := []Player{
		{Name: "Chris", Wins: 10},
		{Name: "Jane", Wins: 5},
	}
	err := json.NewEncoder(w).Encode(leagueTable)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	_, _ = fmt.Fprint(w, score)
}
