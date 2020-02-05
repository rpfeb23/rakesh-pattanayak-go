package main

import "fmt"

const (

	year1 = 2020 + iota   // iota = 0
	year2 = year1 + iota  // iota = 1
	year3 = year1 + iota  // iota = 2
)

func main()  {
	fmt.Println(year1, year2, year3)
}
