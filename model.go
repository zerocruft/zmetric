package zmetric

import "time"

type dataPoint struct {
	timestamp int64
	point     float64
}

type Rate struct {
	Count    int
	Average  float64
	Minimum  float64
	Maximum  float64
	Duration time.Duration
}
