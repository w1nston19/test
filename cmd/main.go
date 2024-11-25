package main

import (
	sleeper2 "github.com/w1nston19/test/sleeper"
	"os"
	"time"
)

func main() {
	sleeper := &sleeper2.ConfigurableSleeper{Duration: 1 * time.Second, SleepF: time.Sleep}
	sleeper2.Countdown(os.Stdout, sleeper)
}
