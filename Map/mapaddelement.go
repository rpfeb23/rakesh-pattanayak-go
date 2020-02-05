package main

import "fmt"

func main()  {
	m := map[string]int{
		"Rakesh" : 85050,
		"Rajesh" : 754005,
	}
	fmt.Println(m)

	m["Adi"] = 85050
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i,v)
	}
}
