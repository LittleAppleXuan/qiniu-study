package main
import (
//	"time"
	"fmt"
	"sync"
)


func main() {

	var wg sync.WaitGroup
	var a int64 = 100000

	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if a > 0 {
					a--
				} else {
					break
				}
			}
		}()
		wg.Wait()
	}

//	time.Sleep(time.Second)
	fmt.Println(a)
}
