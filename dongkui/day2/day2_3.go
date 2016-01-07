package day2
import "errors"

func GetMean(num []float64) (float64,error) {

	var err error
	var sum float64
	var result float64
	for  _, c := range num {
		sum = sum + c
	}
	result = sum / float64(cap(num))
	if cap(num) == 0 {
		errors.New("传入的切片为空!")
	}
	return result,err
}
