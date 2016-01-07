package average

import (
	"errors"
)

//func Division(a,b int) (int,error){
//     if(b==0){
//		 return 0,errors.New("除数不能为0")
//     }
//	return a/b,nil
//}

func Average(a,b float64)(float64,error)  {
	     if(a<0||b<0){
			 return 0,errors.New("a,b不为负数")
	     }
	return (a+b)/2,nil
}

//func main()  {
//	var result float64
//	result,nil=Average(5,4)
//	fmt.Printf(result)
//}