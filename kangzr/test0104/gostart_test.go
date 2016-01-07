package kangzr

import (
	"fmt"
	"testing"
)

func TestModifyAge(t *testing.T) {
	var student Student
	student.Name = "张三"
	student.Age = 22
	student.Height = 180
	student.Number = "20160101"

	student, _ = ModifyAge(33, student)
	if student.Age == 33 {
		t.Log("pass ok")
	} else {
		t.Error("modify failed")
	}
	fmt.Println(student)

}

//func TestModifyAgeFail(t *testing.T){
//	var student Student
//	student.Name="张三"
//	student.Age=22
//	student.Height=180
//	student.Number="20160101"
//
//	student,_=ModifyAge(-1, student)
//	fmt.Println(student)
//	t.Log("ok")
//
//}
