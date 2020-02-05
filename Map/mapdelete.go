package main

import "fmt"

func main()  {
	m := map[string]int{
		"Rakesh" : 85050,
		"Sukanya" : 85050,
		"Rajesh" : 98610,
	}
	fmt.Println(m)
	delete(m,"Sukanya")
	fmt.Println(m)
}
