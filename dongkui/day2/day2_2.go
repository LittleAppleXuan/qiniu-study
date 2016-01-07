package day2
import(
	"fmt"
)

type String string

func (str *String)RevStr() (string,error) {
	newStr := ""
	for _,c := range string(*str) {
//		fmt.Println(c,string(c))
		newStr = string(c) + newStr
	}
	fmt.Println()
	return newStr,nil

//	num := len(*str)
//	mySlice := make([]string,num)
//
//	for i,c := range *str {
//		fmt.Println(c,string(c))
//		mySlice[num-1-i] = string(c)
//	}
//	fmt.Println()
//
//	var newStr string
//	for i,c := range mySlice {
//		fmt.Println(i,c)
//		newStr = newStr + c
//		fmt.Println(newStr)
//	}
//
//	fmt.Println()
//
//	return newStr,nil
}
