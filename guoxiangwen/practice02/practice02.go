package practice02

import (
	"fmt"
	"unsafe"
)

/*
*@ count str length & str byte
 */
func CountStrLenAndByte(str string) (strLen int, strByte uintptr) {
	return len(str), unsafe.Sizeof(str)
}

/**
@ reverse str
*/

func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/**
*@ slice average
 */

func SliceAverage(a []float64, n int) (b float64) {
	var sum float64
	for _, num := range a {
		sum += num
	}
	fmt.Println("sum:", sum)
	b = sum / float64(n)
	return b
}

/**
*@ Println student info order by No
 */
type student struct  {
	id int
	name string
	age int

}

func OrderById() {
	s1:=student{03,"zhangsan",20}
	s2:=student{04,"lisi",23}
	s3:=student{05,"wangwu",24}
	fmt.Println("s1",s1)
	fmt.Println("s2",s2)
	fmt.Println("s3",s3)

	m:=[]student{s1,s2,s3}
	fmt.Println("m",m)




}

/**
* search the index of the string
 */
func IndexOf(orginal,sub string)  {

}
