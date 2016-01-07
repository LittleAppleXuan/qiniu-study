package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetLen(t *testing.T) {
	//中文一个汉字=3个字节
	var strlen, bytelen int = GetLen("12345678中文")
	assert.Equal(t, 10, strlen)
	assert.Equal(t, 14, bytelen)

}
