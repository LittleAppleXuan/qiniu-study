package test0105
import (
	"strings"
	"fmt"
	"errors"
)

//字符串匹配算法
func StrMatch(s1, s2 string) (index []int, err error) {
	if !strings.Contains(s1, s2) {
		err = errors.New("字符串1不包含字符串2!")
		return nil, err
	}

	count := int(strings.Count(s1, s2))
	if count > 0 {
		fmt.Println(s1, "包含", s2, "的位置个数有:", count)

		inFirst := strings.Index(s1, s2)
		index = append(index, inFirst)
		count = count-1

		inLast := strings.LastIndex(s1, s2)
		if inLast != inFirst {
			index = append(index, inLast)
			count=count-1

		}

		return index, nil
	} else {
		return index, err
	}
}
