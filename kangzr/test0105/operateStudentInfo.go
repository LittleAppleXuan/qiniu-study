package test0105
import (
	"fmt"
	"sort"
)

type Student struct {
	Number int
	Name   string
	Age    int
}

//保存多个学生信息,并按学号排序打印出来
func OperateStudentInfo(students []Student) {
	var length int
	length = int(len(students))
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = students[i].Number
		//		fmt.Println(numbers[i])
	}
	fmt.Println("学号数组:", numbers)
	sort.Ints(numbers)
	fmt.Println("排序后的学号数组", numbers)

	var studentOut []Student
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if numbers[i] == students[j].Number {
				studentOut = append(studentOut, students[j])
			}
		}
	}

	fmt.Println("按学号排序后的学生信息输出为:")
	fmt.Println(studentOut)
}
