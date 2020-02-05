// sort Package has an Interface. if you code to that  Interface you can use any of the sort functions that takes sort Interface as their argument https://golang.org/pkg/sort/#Interface

package main

import (
	"fmt"
	"sort"
)

type person struct{
	name string
	age int
}

type sortpersonbyage []person

// Implement below 3 functions to code to Sort Interface

func (s sortpersonbyage) Len() int {
	return len(s)
}
// On what basis you will sort and Increasing or decreasing
// Return True means Don't Swap (Items are SORTED do nothing)
// Return False Means Swap (Items are NOT Sorted)

func (s sortpersonbyage) Less(i, j int) bool {
	if s[i].age > s[j].age{
		return false
	}else {
		return true
	}
}

func (s sortpersonbyage) Swap(i, j int)  {
    s[i], s[j] = s[j], s[i]
}

func main()  {

	p1 := person{
		name: "Rakesh",
		age:  35,
	}
	p2 := person{
		name: "James Bond",
		age:  25,
	}
	p3 := person{
		name: "Mis Moneypenny",
		age:  45,
	}

	// SORT BY AGE
    p := sortpersonbyage{p1,p2,p3}  // sortpersonbyage is of type data Interface in sort package as it implements the 3 functions
    fmt.Println("before sort : ",p)
    sort.Sort(p)
    fmt.Println("After sort",p)

}