package main
import (
	"github.com/qiniu-study/dongkui/day2"
	"log"
	"fmt"
)

func main() {
	//2.1
	charNum,byteNum,err := day2.CountString("Hello,world and 程序员!")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("charNum:",charNum,",byteNum:",byteNum)

	//2.2
	var str day2.String
	str = "Hello ,world and 程序员!"
	newStr,err := str.RevStr()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("newStr:",newStr)

	//2.3
	mySlice := []float64{11.11,22.22,33.33,44.44}
	mean,_ := day2.GetMean(mySlice)
	fmt.Println("平均值:",mean)

	//2.4

//	student1 := day2.Student{Name:"a",Age:18,Num:"3"}
//	student1.SaveInf()
//	fmt.Println(day2.Students)
//
//	student2 := day2.Student{"b",18,"2"}
//	student2.SaveInf()
//	fmt.Println(day2.Students)
//
//	student3 := day2.Student{"c",18,"1"}
//	student3.SaveInf()



}
