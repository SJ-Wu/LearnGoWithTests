package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store: store}
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
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

func (p *PlayerServer) GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "Floyd" {
		return "10"
	}
	return "0"
}

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) RecordWin(player string) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	player = ""
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
