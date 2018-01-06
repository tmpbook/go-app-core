package Stopwatch

import (
	"fmt"
	"testing"
	"time"
)

func TestEndStopwatch(t *testing.T) {
	sw := NewStopwatch()

	<-time.After(50 * time.Millisecond)
	fmt.Printf("Cost: %d ms\n", sw.Lap("once"))

	<-time.After(2000 * time.Millisecond)
	fmt.Printf("Cost: %d ms\n", sw.Lap("once"))

	<-time.After(500 * time.Millisecond)
	fmt.Printf("Cost: %d ms\n", sw.End())

	sw.PrintAll()
}
