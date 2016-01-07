package person
import (
	"fmt"
	"study/person/person"
)

func day1() {
	fmt.Println("main func")
	var person person.Person
	person.SetPerson("xiaoming","19","man","178","65")
	person.SayHello()
	person.ShowInf()
	person.SayPlusTable()
}
