package main

import "fmt"

func main()  {
	var x = make([]int,5, 10)
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	x[4] = 42
	x[0] = 99

	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	// You cant do below though capacity. Index Out of Range

	//x[5] = 55

	x = append(x,55,66,77,88,99)

	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	// This will extend beyond its capacity
    x = append(x, 11)

	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)
}
