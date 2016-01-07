package kangzr

import "fmt"

//func main() {
//	fmt.Println("hello world")
//
//	var student Student
//	student.Name="张三"
//	student.Age=22
//	student.Height=180
//	student.Number="20160101"
//
////	fmt.Println("修改前的年龄为:", student.Age)
//	student,_=ModifyAge(25, student)
////	fmt.Println("修改后的年龄为:", student.Age)
//	fmt.Println(student)
//
//
//}

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Number string `json:"number"`
}

func ModifyAge(newAge int, student Student) (st Student, err error) {
	fmt.Println("修改前的年龄为:", student.Age)
	student.Age = newAge
	fmt.Println("修改后的年龄为:", student.Age)
	st = student
	return st, err
}
