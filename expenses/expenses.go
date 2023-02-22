package expenses

func Avegare(expns ...float32) float32 {
	return Sum(expns...) / float32(len(expns))
}
func Sum(expns ...float32) float32 {
	var sum float32

	for _, e := range expns {
		sum += e
	}

	return sum
}
func Max(expns ...float32) float32 {
	var max float32

	for _, e := range expns {
		if e > max {
			max = e
		}
	}

	return max
}
func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}

	var min float32 = expns[0]

	for _, e := range expns {
		if min > e {
			min = e
		}
	}

	return min
}
