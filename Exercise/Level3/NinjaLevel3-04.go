package main

import "fmt"

func main()  {
	my_birth_year := 1985

	for {
		if my_birth_year > 2020 {
			break
		}
		fmt.Println(my_birth_year)
		my_birth_year++
	}
}
