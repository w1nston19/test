package main

import (
	"os"
	s "test/test/sleeper"
	"time"
)

func main() {
	sleeper := &s.ConfigurableSleeper{1 * time.Second, time.Sleep}
	s.Countdown(os.Stdout, sleeper)
}
