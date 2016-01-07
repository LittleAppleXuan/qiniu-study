package main

import (
	// student "./student"
	stringutils "./../stringutils"
	"errors"
	"log"
)

func average(numbers []float64) (float64, error) {

	if len(numbers) == 0 {
		return 0, errors.New("param not be null")
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	log.Println(numbers)
	log.Println(sum)

	return sum / float64(len(numbers)), nil

}
func main() {
	// numbers := make([]float64, 0)
	// numbers = append(numbers, 1.234)
	// numbers = append(numbers, 1.345)
	// numbers = append(numbers, 1.456)
	// a, err := average(numbers)

	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Printf("average is %f", a)
	// xueliu := student.Student{"2005", "xueliu", 19}
	// wangwu := student.Student{"1010", "wangwu", 12}
	// lisi := student.Student{"1003", "lisi", 14}
	// zhangsan := student.Student{"1005", "zhangsan", 17}

	// student.Add(wangwu)
	// student.Add(lisi)
	// student.Add(zhangsan)
	// student.Add(xueliu)

	// log.Println(student.GetStudents())

	re := make([]int, 0)
	stringutils.FindSubStrIndex("rtuyyuyuxxxuc", "u", &re)
	log.Println("result", re)

}
