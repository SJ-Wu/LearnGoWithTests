package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const countdownStart = 3
const finalWord = "Go!"

func Countdown(write io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, err := fmt.Fprintln(write, i)
		if err != nil {
			return
		}
	}
	sleeper.Sleep()
	_, err := fmt.Fprintln(write, finalWord)
	if err != nil {
		return
	}
}
