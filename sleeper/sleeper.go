package sleeper

import (
	"fmt"
	"io"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep() error
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	Duration time.Duration
	SleepF   func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() error {
	c.SleepF(c.Duration)
	return nil
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
