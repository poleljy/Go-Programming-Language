package ch7

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func TestSleep() {
	flag.Parse()

	fmt.Printf("Sleep for %v...", *period)
	time.Sleep(*period)
}
