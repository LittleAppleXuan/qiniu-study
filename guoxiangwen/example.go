package main

import (
	"fmt"
	"github.com/guoxiangwen/golang-train/add"
	"qiniu-study/guoxiangwen/practice02"
	"strings"
)

func main() {
	s, err := add.Add(4, 4)
	if err != nil {
		fmt.Printf("error")
		return
	}
	fmt.Println("the plus is:", s)

	strLen, strByte := practice02.CountStrLenAndByte("hahaha")
	fmt.Println("the str length is:", strLen)
	fmt.Println("the str byte is:", strByte)

	a := "Hello"
	println(a)
	println(practice02.Reverse(a))

	myArray := make([]float64, 3)
	myArray[0] = 1.234
	myArray[1] = 1.344
	myArray[2] = 1.454
	avg := practice02.SliceAverage(myArray, len(myArray))
	fmt.Println("avg", avg)
	index:=strings.Index("qwerter","er")
	fmt.Println("index",index)


	practice02.OrderById()
}
