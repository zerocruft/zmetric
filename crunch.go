package zmetric

func crunchRate(points []dataPoint) Rate {
	rate := Rate{}

	var sum int64 = 0
	var min = points[0].point
	var max = points[0].point

	for _, dp := range points {
		sum += dp.point

		if dp.point < min {
			min = dp.point
		}
		if dp.point > max {
			max = dp.point
		}
	}
	rate.Average = sum / int64(len(points))
	rate.Maximum = max
	rate.Minimum = min
	rate.Count = len(points)

	return rate
}
