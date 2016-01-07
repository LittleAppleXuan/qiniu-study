package stringutils

import (
	"strings"
	"unicode/utf8"
)

func GetLen(str string) (int, int) {

	return utf8.RuneCountInString(str), len(str)

}

func FindSubStrIndex(str, subStr string, result *[]int) {

	first := strings.Index(str, subStr)

	strs := strings.Split(str, "")
	strslice := strs[first+1:]
	strs1 := strings.Join(strslice, "")
	if len(*result) > 0 {
		first = (*result)[len(*result)-1] + first + 1
	}
	*result = append(*result, first)
	if strings.Index(strs1, subStr) != -1 {
		FindSubStrIndex(strs1, subStr, result)
	}

}
