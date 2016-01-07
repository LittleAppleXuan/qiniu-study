package day2
import "unsafe"

func CountString(inStr string) (int,int,error){
	charNum := len(inStr)
	var byteNum = unsafe.Sizeof(inStr)
	return charNum,int(byteNum),nil
}
