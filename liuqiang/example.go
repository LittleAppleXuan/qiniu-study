package main

import (
	"fmt"
	"./average"
)

func main()  {
	result,err:=average.Average(-2,3)
	if result!=0.5||err!=nil{
		fmt.Println("测试数据不能为负数")
		return
	}
	fmt.Println("The average is",result)
}