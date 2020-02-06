// Use go run -race racecondition.go to see if race condition exist
package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)
var wg sync.WaitGroup
var x int64
func main()  {

	wg.Add(2)
	go increment1()
	go increment2()
	wg.Wait()
}

func increment1()  {

	for i := 0; i <100 ; i++{
		atomic.AddInt64(&x,1)
		fmt.Println("Increment1 : ", atomic.LoadInt64(&x))
	}
	wg.Done()
}

func increment2()  {
	for i := 0; i <100 ; i++{
		atomic.AddInt64(&x,1)
		fmt.Println("Increment2 : ", atomic.LoadInt64(&x))
	}
	wg.Done()
}