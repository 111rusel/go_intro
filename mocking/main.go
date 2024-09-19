package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func main() {
	s := &DefaultSleeper{}
	CountDown(os.Stdout, s)
}

func CountDown(out io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, finalWord)
	// for i := countDownStart; i > 0; i-- {

	// }
}
