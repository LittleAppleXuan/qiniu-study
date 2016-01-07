package average_test

import (
	"testing"
)

//func Test_Division(t *testing.T){
//	if i,e:=Division(6,2);i!=3||e!=nil {
//		t.Error("不通过")
//	}else {
////		t.Log("通过")
//		fmt.Println("通过")
//	}
//}

func Test_Average1(t *testing.T)  {
	if i,e:=Average(2.0,4.0);i!=3.0||e!=nil {
		t.Error("不通过")
	}else {
		t.Log("通过")
	}
}

func Test_Average2(t *testing.T)  {
	if i,e:=Average(3.0,4.0);i!=3.0||e!=nil {
		t.Error("不通过")
	}else {
		t.Log("通过")
	}
}

func Test_Average3(t *testing.T)  {
	if i,e:=Average(-2,4);i==0||e!=nil {
		t.Error(e)
	}else {
		t.Log("通过")
	}
}