package test0105

//反转字符串
func ReverseString(str string) string {
	newStr := ""
	for _, c := range str {
		newStr = string(c) + newStr
	}
	return newStr
}