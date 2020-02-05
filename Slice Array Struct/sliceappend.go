package main

import "fmt"

func main(){
	var x = []int{1,2,3,4,5}
	fmt.Println(x)
	var y = []int{91,92,93,94,95}

	x = append(x,11,12,13)
	fmt.Println(x)

	// dumb way to do is range. check the next logic for x and z append
	for _,v := range y{
		x = append(x,v)
		fmt.Println(x)
	}

	var z = []int{111, 222, 333, 444}
	x = append(x, z...) // z... is NOT Variadic Parameter. It is unfurling elements of z. Append function declartion has variadic parameter of same slice type func append(s []T, vs ...T) []T

	fmt.Println(x)
}
