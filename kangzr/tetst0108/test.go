package main
import "fmt"


func main()  {
	var s string
//	s = "hello"

	if len(s) == 0 {
		fmt.Println("nil")
	} else {
		fmt.Println(len(s))
	}
}
