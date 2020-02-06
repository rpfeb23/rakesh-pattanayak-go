// Use go run -race racecondition.go to see if race condition exist
package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var mu sync.Mutex
var x int
func main()  {

	wg.Add(2)
	go increment1()
	go increment2()
	wg.Wait()
}

func increment1()  {

	for i := 0; i <100 ; i++{
		mu.Lock()
		x++
		fmt.Println("Increment1 : ", x)
		mu.Unlock()
	}
    wg.Done()
}

func increment2()  {
	for i := 0; i <100 ; i++{
		mu.Lock()
		x++
		fmt.Println("Increment2 : ", x)
		mu.Unlock()
	}
	wg.Done()
}