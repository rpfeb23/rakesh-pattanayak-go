// Slice is built on top of array. When Slice grows, a new array is created and data gets stored. Old array is thrown away. This takes some processing power.
// Slice Capacity is NOTHING but Underlying Array Size
package main

import "fmt"

func main()  {
	var x []int
	fmt.Printf("Slice x : length : %v  Capacity : %d  %v\n",len(x),cap(x),x)

	x = []int{10,20,30}
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	// Look At How CAPACITY changes
	// When you are out of Capacity, It is DOUBLED
	fmt.Println("Capacity will be DOUBLED. Observe")
    x = append(x,40)
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	x = append(x,50)
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	x = append(x,60)
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)

	// When you are out of Capacity, It is DOUBLED
	fmt.Println("Capacity will be DOUBLED AGAIN")
	x = append(x,70)
	fmt.Printf("Slice x : length : %v  Capacity : %v  %v\n",len(x),cap(x),x)
}
