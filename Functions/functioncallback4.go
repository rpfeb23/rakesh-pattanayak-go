package main

import "fmt"

func main()  {
    main_even_sum := evensum(sum,1,2,3,4,5,6,7,8)
    fmt.Println(main_even_sum)
    main_odd_sum := oddsum(sum, 1,2,3,4,5,6,7,8)
    fmt.Println(main_odd_sum)
}

func sum(x ...int) int{
	sum :=0
	for _, v := range x{
		sum = sum + v
	}
	return sum
}
func evensum(f func(x ...int) int, y ...int)  int{
	evennums := []int{}
	for _,v := range y{
		if v % 2 == 0{
			evennums = append(evennums,v)
		}
	}
	fmt.Println("List of even numbers are : ", evennums)
	evensum := f(evennums...)
	return evensum
}

func oddsum(f func(x ...int) int, y ...int) int {
	oddnumbers := []int{}
	for _, v := range y{
		if v %2 != 0 {
			oddnumbers = append(oddnumbers, v)
		}
	}
	fmt.Println("List of Odd Numbers : ", oddnumbers)
	return (f(oddnumbers...))
}
