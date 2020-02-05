package DatatypesLoops

import "fmt"

func main()  {
	for i := 33; i <= 122 ; i++ {
		fmt.Printf("Decimal : %d Character : %v  UTF-8 with ASCII : %#U\n", i, string(i), i)
	}
}
