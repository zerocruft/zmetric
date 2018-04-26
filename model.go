package zmetric

import (
	"encoding/json"
	"time"
)

type dataPoint struct {
	timestamp int64
	point     int64
}

type Rate struct {
	Count    int
	Average  int64
	Minimum  int64
	Maximum  int64
	Duration time.Duration
}

func (r Rate) String() string {
	rb, _ := json.Marshal(r)
	return string(rb)
}
