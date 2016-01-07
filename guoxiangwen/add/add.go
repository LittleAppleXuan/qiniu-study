package add

import (
	"fmt"
)

func Add(a, b int) (int, error) {
	fmt.Printf("begin plus")
	return a + b, nil
}
