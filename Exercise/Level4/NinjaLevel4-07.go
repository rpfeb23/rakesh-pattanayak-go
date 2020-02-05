package main

import "fmt"

func main()  {
	var multidimensionslice [][]string

	multidimensionslice = [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Println(multidimensionslice)

	for i, v:= range multidimensionslice{
		fmt.Println(i,v)
	}
    fmt.Println("************************************")
	for i, v:= range multidimensionslice{
		fmt.Println(i,v)
		for i1, v1:= range v{
			fmt.Println("                   ",i1,v1)
		}
	}
}


