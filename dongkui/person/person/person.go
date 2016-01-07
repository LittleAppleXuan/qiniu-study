package person

import (
	"fmt"
)

type Person struct {
	Name string
	Age string
	Sex string
	Height string
	Weight string
}

func (person *Person) SetPerson(name, age, sex, height, weight string) error {
	person.Name = name
	person.Age = age
	person.Sex = sex
	person.Height = height
	person.Weight = weight
	return nil
}

func (person *Person) SayHello() error {
	fmt.Println("Hello!Nice to see you.My name is "+person.Name)
	return nil
}

func (person *Person) ShowInf() error {
	fmt.Println(person)
	return nil
}

func (person *Person) SayPlusTable() error {
	for i := 1;i < 10;i++{
		for j := 1;j <= i;j++{
			fmt.Print(i,"*",j,"=",i*j,",")
		}
		fmt.Println("")
	}
	return nil
}