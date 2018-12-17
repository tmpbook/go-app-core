package stopwatch

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	// STARTNAME is name of the time which Stopwatch start from
	STARTNAME = "0-Start"
	// ENDNAME is name of the time which Stopwatch end in
	ENDNAME = "(LastLap)"
)

type lap struct {
	name string
	time time.Time
}

// Stopwatch struct record the lap name and its cost durations(millisecond)
type Stopwatch struct {
	laps []lap
	mu   sync.Mutex
}

// NewStopwatch create a Stopwatch and record time with the name of STARTNAME
// Return a point of own
func NewStopwatch() *Stopwatch {
	sw := new(Stopwatch)
	sw.laps = append(sw.laps, lap{STARTNAME, time.Now()})
	return sw
}

func (sw *Stopwatch) recordOneLapWithName(name string) uint32 {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	// get the latest lap's time
	latestLap := sw.laps[len(sw.laps)-1]
	now := time.Now()
	sw.laps = append(sw.laps, lap{name, now})
	return uint32(now.Sub(latestLap.time) / time.Millisecond)
}

// Lap the time duration from lastest lap
func (sw *Stopwatch) Lap(name string) uint32 {
	// index the lap name
	lapName := `(Lap` + strconv.Itoa(len(sw.laps)) + `)` + name
	return sw.recordOneLapWithName(lapName)
}

// End function don't need a name
func (sw *Stopwatch) End() uint32 {
	return sw.recordOneLapWithName(ENDNAME)
}

// PrintAll of Stopwatch status now
func (sw *Stopwatch) PrintAll() {
	for i := 1; i < len(sw.laps); i++ {
		this := sw.laps[i].time
		lastest := sw.laps[i-1].time
		fmt.Printf("%s cost: %d ms;\n", sw.laps[i].name, this.Sub(lastest)/time.Millisecond)
	}
}
