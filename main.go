package main

import (
	"os"
	sleeper2 "test/sleeper"
	"time"
)

func main() {
	sleeper := &sleeper2.ConfigurableSleeper{1 * time.Second, time.Sleep}
	sleeper2.Countdown(os.Stdout, sleeper)
}
