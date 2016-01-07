package chapter2

import (
	"fmt"
	"sort"
)

type Student struct {
	Id   string
	Name string
	Age  int
}

var students map[string]Student
var keys []string

func init() {
	students = make(map[string]Student)
}
func Add(student Student) {
	students[student.Id] = student
	keys = append(keys, student.Id)

}
func GetStudents() map[string]Student {

	sort.Strings(keys)
	sortStudents := make(map[string]Student)
	for _, k := range keys {
		sortStudents[k] = students[k]
	}

	return sortStudents
}
