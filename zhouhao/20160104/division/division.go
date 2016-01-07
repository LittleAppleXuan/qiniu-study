package division
import (
	"errors"
)
func Division(dividend,divisor int)(int,error)  {
	if(divisor==0){
		return 0,errors.New("除数只能非0!!!")
	}
	return dividend/divisor,nil
}

