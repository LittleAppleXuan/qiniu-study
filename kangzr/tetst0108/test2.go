package main
import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
	"time"
)


func main() {

	var a int64 = 100000

	_ = sync.Mutex{}
	_ = atomic.AddInt64

	for n := 0; n < 10; n++ {
		go func() {
			for {
				if a > 0 {
					atomic.AddInt64(&a, -1)
					runtime.Gosched()  // 允许其他协程进行处理
				} else {
					break
				}
			}
		}()
	}

	time.Sleep(time.Second)
	opsA := atomic.LoadInt64(&a)

	fmt.Println(opsA)
}
