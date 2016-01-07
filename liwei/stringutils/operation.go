package stringutils

import (
	"strings"
)

func Invert(str string) string {
	strs := strings.Split(str, "")
	slen, _ := GetLen(str)
	invert := make([]string, slen)
	for i, item := range strs {
		invert[slen-1-i] = item

	}
	return strings.Join(invert, "")

}
