package timer

import (
	"os"
	"fmt"
	"time"
	"math"
)

// Countdown runs a basic time left countdown and shows time left since start
func Countdown(target time.Time) {
	fmt.Println("🍅 Starting gomodoro timer ...")

	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", FormatMins(0), "  \r ")
			return
		}
		fmt.Fprint(os.Stdout, "Countdown: ", FormatMins(timeLeft), "   \r ")
		os.Stdout.Sync()
	}
	fmt.Print("hello")
}

// FormatSecs returns proper time left by secs
func FormatSecs(timeLeft time.Duration) string {
	return fmt.Sprintf("%02.1f", math.Abs(timeLeft.Seconds()))
}

// FormatMins returns proper time left by mins
func FormatMins(timeLeft time.Duration) string {
	minutes := int(timeLeft.Minutes())
	seconds := int(timeLeft.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}
