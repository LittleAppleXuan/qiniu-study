package person
import (
	"testing"
)

func TestSetPerson(t *testing.T) {
	var person Person
	person.SetPerson("xiaoming","19","man","178","65")
	if person.Name == "xiaoming" {
		t.Log("success")
	}else {
		t.Error("failed")
	}
}
