package zmetric

import (
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

var (
	rates      map[string]Rate
	ratesMutex sync.Mutex
)

func init() {
	rates = map[string]Rate{}
	ratesMutex = sync.Mutex{}
}

// NewRate creates a float64 channel that is monitored with at least the given sample size time.Duration passed in.
// The channel is used to send float64 data point values and the Get function will return the crunched Rate object
// created from monitoring subject channel
func NewRate(sampleDuration time.Duration) (chan int64, string, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return nil, "", err
	}
	key := id.String()

	gauge := make(chan int64)
	points := []dataPoint{}
	pointsMutex := sync.Mutex{}

	go func() {
		for {
			point := <-gauge
			dp := dataPoint{
				point:     point,
				timestamp: time.Now().UnixNano(),
			}
			points = append(points, dp)
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for t := range ticker.C {
			timeRate := t.Add(-1 * sampleDuration).UnixNano()
			newPoints := []dataPoint{}
			pointsMutex.Lock()
			for _, dp := range points {
				if dp.timestamp > timeRate {
					newPoints = append(newPoints, dp)
				}
			}
			points = newPoints
			pointsMutex.Unlock()

			go func() {
				rate := crunchRate(points)
				rate.Duration = sampleDuration
				ratesMutex.Lock()
				rates[key] = rate
				ratesMutex.Unlock()
			}()
		}
	}()

	return gauge, key, nil
}

// Get will return the Rate associated with the given key, or nil if no Rate is associated with given key
func Get(key string) Rate {
	return rates[key]
}
