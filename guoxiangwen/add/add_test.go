package add

import (
	"testing"
)

func Test_Add_1(t *testing.T) {
	if i, e := Add(5, 4); i != 9 || e != nil {
		t.Error("the result is not right")
	} else {
		t.Log("pass ok")
	}
}
