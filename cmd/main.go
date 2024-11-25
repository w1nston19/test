package main

import (
	sleeper2 "github.com/w1nston19/test/sleeper"
	"os"
	"time"
)

func main() {
	sleeper := &sleeper2.ConfigurableSleeper{1 * time.Second, time.Sleep}
	sleeper2.Countdown(os.Stdout, sleeper)
}
