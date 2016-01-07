package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Invert(t *testing.T) {
	assert.Equal(t, "国中rts", Invert("str中国"))
}
