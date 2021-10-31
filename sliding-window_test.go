package slidingwindow

import (
	"fmt"
	"testing"
	"time"
)

type Record struct {
	Conductivity float64
	Temperature  float64
	Depth        float64
}

func process(w *SlidingWindow) (*Stats, *Stats, *Stats) {
	records := w.Get()
	cStats := NewStats()
	tStats := NewStats()
	dStats := NewStats()
	for _, record := range records {
		if r, ok := record.(*Record); ok {
			cStats.Append(r.Conductivity)
			tStats.Append(r.Temperature)
			dStats.Append(r.Depth)
		}
	}
	return cStats, tStats, dStats
}

func TestSlidingWindow(t *testing.T) {
	w := NewSlidingWindow(5 * time.Second)
	cStats, tStats, dStats := process(w)
	fmt.Printf("c: %v, t: %v, d: %v\n", cStats, tStats, dStats)

	w.Append(&Record{
		Conductivity: 10,
		Temperature:  100,
		Depth:        1000,
	})
	cStats, tStats, dStats = process(w)
	fmt.Printf("c: %v, t: %v, d: %v\n", cStats, tStats, dStats)

	time.Sleep(2 * time.Second)
	w.Append(&Record{
		Conductivity: 20,
		Temperature:  200,
		Depth:        2000,
	})
	cStats, tStats, dStats = process(w)
	fmt.Printf("c: %v, t: %v, d: %v\n", cStats, tStats, dStats)

	time.Sleep(2 * time.Second)
	w.Append(&Record{
		Conductivity: 30,
		Temperature:  300,
		Depth:        3000,
	})
	cStats, tStats, dStats = process(w)
	fmt.Printf("c: %v, t: %v, d: %v\n", cStats, tStats, dStats)

	time.Sleep(2 * time.Second)
	w.Append(&Record{
		Conductivity: 40,
		Temperature:  400,
		Depth:        4000,
	})
	cStats, tStats, dStats = process(w)
	fmt.Printf("c: %v, t: %v, d: %v\n", cStats, tStats, dStats)
}
