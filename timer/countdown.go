package timer

import (
	"os"
	"fmt"
	"time"
)

// Countdown indicates timeleft since the timer started
func countdown(target time.Time) {
	fmt.Println("Starting timer ...")

	for range time.Tick(1000 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Printf("Finished timer %v", 0)
			return
		}
		timeLeft = time.Duration(timeLeft)
		fmt.Fprint(os.Stdout, "#: ", formatSeconds(timeLeft))
	}
	return
}

// formatSeconds formats the time durations in seconds
func formatSeconds(t time.Duration) time.Duration {
	//return time.Duration(t.Seconds()) * 1000
	return time.Duration(int64(time.Millisecond) * int64(time.Duration(t.Seconds())))
}
