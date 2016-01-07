package test0105

import (
	"unsafe"
)
//统计字符串的字符数与字节数
func GetCharNumber(str string) (charNumber, byteNumber int) {
	charNumber = len(str)
	byteNumber = int(unsafe.Sizeof(str))
	return charNumber, byteNumber
}
