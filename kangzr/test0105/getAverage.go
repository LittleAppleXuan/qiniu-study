package test0105

//求slice的平均值
func CalculateAvrage(slice []float64) float64 {
	var totalFloat float64
	totalFloat = 0.0
	for i := 0; i < len(slice); i++ {
		totalFloat = totalFloat + slice[i]
	}
	average := float64(totalFloat / float64(len(slice)))
	return average
}
