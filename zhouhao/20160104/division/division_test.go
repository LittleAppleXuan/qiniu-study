package division
import (
	"testing"
	"fmt"
)

func TestDivision1(t *testing.T)  {
	const dividend,divisor,result= 8,4,2
	if r,e:=Division(dividend,divisor);r!=result||e!=nil{
		t.Errorf("测试不通过!!!")
	}else {
		fmt.Println("ok")
//		t.Log("通过")
	}
}
func TestDivision2(t *testing.T)  {
	const dividend,divisor,result= 8,0,2
	if r,e:=Division(dividend,divisor);r!=result||e!=nil{
		t.Errorf("测试不通过!!!")
	}else {
		fmt.Println("ok")
		//		t.Log("通过")
	}
}