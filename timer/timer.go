package timer

import (
	"fmt"
	"time"
)

// GomodoroTimer defines the structure of a pomodoro timer
type GomodoroTimer struct {
	Start         time.Time
	Duration 	  time.Duration
	MaxIntervals  int
	ShortBreak    time.Duration
	LongBreak     time.Duration
}

// SetTimer runs the pomodoro timer
func SetTimer (timer GomodoroTimer) {
	for i := 1; i < timer.MaxIntervals; i++ {

		// initialize start time
		start := time.Now()

		// calculate finish time by adding focus duration
		finish := start.Add(timer.Duration)

		fmt.Printf("Begin interval %v focus time at %v. \n", i, start)

		// count down until end of focus time
		Countdown(finish)

		// timer end message
		fmt.Printf("End of interval %v focus time at %v. \n", i, finish)
		fmt.Println("5 minute break!")

		// count down until end of break time
		breaktime := finish.Add(timer.ShortBreak)
		Countdown(breaktime)
	}
}

