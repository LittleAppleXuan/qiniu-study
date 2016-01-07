package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsEmail(t *testing.T) {
	assert.True(t, IsEmail("contack@gmail.com"))

}

func Test_IsEmail2(t *testing.T) {
	assert.False(t, IsEmail("contackidrmfly.com"))
}

func Test_IsPhone(t *testing.T) {
	assert.True(t, IsPhone("18502842115"), IsPhone("18533333333"), IsPhone("17000000000"), IsPhone("8618981125210"))

}
func Test_IsPhone2(t *testing.T) {
	assert.False(t, IsPhone("123"), IsPhone("7613333333333"), IsPhone("14555555555"))

}
func Test_IsIDCard(t *testing.T) {
	assert.True(t, IsIDCard("620422199901038788"))

}
func Test_IsIDCard2(t *testing.T) {
	assert.False(t, IsIDCard("62042219993003878"))

}
