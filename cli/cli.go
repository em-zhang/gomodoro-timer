package main 

import(
	"time"
	"flag"
	timer "github.com/em-zhang/gomodoro/timer"
)

// create a new pomodoro timer
var pomo = timer.Timer{
	Start:         time.Now(),
	Duration: 	   25 * time.Minute,
	MaxIntervals:  4,
	ShortBreak:    5 * time.Minute,
}

func init() {
	// optional command line params
	duration := flag.Int("duration", 25, "Change the pomodoro time – the default is 25 minutes!")
	shortbreak := flag.Int("shortbreak", 5, "Change the shortbreak time – the default is 5 minutes!")
	intervals := flag.Int("intervals", 5, "Change the number of pomodoro intervals – the default is 5!")

	flag.Parse()
	// update variables as needed
	if *duration != 0 {
		pomo.Duration = time.Duration(*duration) * time.Minute
	}
	if *shortbreak != 0 {
		pomo.ShortBreak = time.Duration(*shortbreak) * time.Minute
	}
	if *intervals != 0 {
		pomo.MaxIntervals = *intervals
	}
}

func main() {
	timer.SetTimer(pomo)
}