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
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	_, _ = fmt.Fprint(w, p.store.GetPlayerScore(player))

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

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
