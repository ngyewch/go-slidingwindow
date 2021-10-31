package slidingwindow

import "math"

type Stats struct {
	Mean  float64
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func NewStats() *Stats {
	return &Stats{
		Mean:  math.NaN(),
		Min:   math.NaN(),
		Max:   math.NaN(),
		Sum:   0,
		Count: 0,
	}
}

func (stats *Stats) Append(value float64) {
	stats.Sum += value
	stats.Count += 1
	stats.Mean = stats.Sum / float64(stats.Count)
	if math.IsNaN(stats.Min) || value < stats.Min {
		stats.Min = value
	}
	if math.IsNaN(stats.Max) || value > stats.Max {
		stats.Max = value
	}
}
