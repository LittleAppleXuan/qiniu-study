package chapter2


type Student struct {
	Id string
	Name string
	Age string
}

var Students := make(map["string"]Student,0)
func Add(student Student) {
     Students = append(Students,student)
	
}