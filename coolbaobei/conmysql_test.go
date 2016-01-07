package db

import (
	"testing"
)

func Test_FindByPk(t *testing.T) {
	num := FindByPone("18080260070")
	t.Log(num)
}
