package main
import "fmt"
import "./firstfile"

func main()  {
	i,e := firstfile.Cheng(8,7)
	fmt.Printf("the result is : %d\n",i)
	fmt.Printf("the err is : %s\n",e)
}
