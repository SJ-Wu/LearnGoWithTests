package main

import (
	"LearnGoWithTests/di"
	"LearnGoWithTests/mocking"
	"net/http"
	"os"
)

func main() {
	countdown()
}

func countdown() {
	mocking.Countdown(os.Stdout, nil)
}

func greet() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(di.MyGreeterHandler))
	if err != nil {
		return
	}
}
