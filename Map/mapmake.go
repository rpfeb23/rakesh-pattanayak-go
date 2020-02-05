package main

import "fmt"

func main()  {
	m := map[string]int{"Rakesh": 85050,
		"Rajesh": 754005,
	}
	fmt.Println(m)

	//var m1 map[string]int
	//fmt.Println(m1)

	// The make function returns a map of the given type, initialized and ready for use.

	var m1 = make(map[string]int)
	fmt.Println(m1)
}
