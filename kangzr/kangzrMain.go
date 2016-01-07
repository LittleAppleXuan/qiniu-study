package main

import (
	test"./test0105"
	"fmt"
)

func main() {
	str := "Hello,世界!"
	charNumber, byteNumber := test.GetCharNumber(str)
	fmt.Println("字符串", str, "的字符数为:", charNumber)
	fmt.Println("字符串", str, "的字节数为:", byteNumber)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("反转前的字符串:", str)
	fmt.Println("反转后的字符串:", test.ReverseString(str))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	var myArray [20]float64 = [20]float64{1.1, 2.3, 3.5, 4.7, 4.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	var mySlice []float64 = myArray[:5]
	fmt.Println("mySlice为:", mySlice)
	fmt.Println("mySlice的平均值为:", test.CalculateAvrage(mySlice))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	var students []test.Student
	student1 := test.Student{34, "李雷", 22}
	student2 := test.Student{21, "韩梅梅", 20}
	student3 := test.Student{65, "张三", 33}
	students = append(students, student1)
	students = append(students, student2)
	students = append(students, student3)
	fmt.Println("保存学生信息为:")
	fmt.Println(students)
	test.OperateStudentInfo(students)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//字符串匹配算法
	s1 := "sdsdsdsdsdsdsdsddfdfdffgf"
	s2 := "ds"
	index, err := test.StrMatch(s1, s2)
	if err != nil {
		panic(err)
	}
	fmt.Println(s1, "包含", s2, "位置有:", index)

}
