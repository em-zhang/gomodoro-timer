package timer

import (
	"fmt"
	"time"
)

// Pomodoro defines the structure of a pomodoro timer
type Timer struct {
	Start         time.Time
	Duration time.Duration
	MaxIntervals  int
	ShortBreak    time.Duration
	LongBreak     time.Duration
}

// SetTimer runs the pomodoro timer
func SetTimer (timer Timer) {
	for i := 1; i < timer.MaxIntervals; i++ {

		// initialize start time
		begin := time.Now()

		// calculate finish time by adding focus duration
		end := begin.Add(timer.Duration)

		fmt.Printf("Begin interval %v focus time %v. \n", i, begin)

		// count down until end of focus time
		countdown(end)

		// timer end message
		fmt.Printf("End of interval %v focus time %v. \n", i, end)
		fmt.Println("5 minute break!")

		// count down until end of break time
		breaktime := end.Add(timer.ShortBreak)
		countdown(breaktime)
	}
}

