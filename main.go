package main

import (
	"LearnGoWithTests/di"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(di.MyGreeterHandler))
	if err != nil {
		return
	}
}
