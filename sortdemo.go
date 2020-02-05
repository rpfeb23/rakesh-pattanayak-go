package main

import (
	"fmt"
	"sort"
)

func main()  {
	sliceint := []int{4,3,1,2,5}
	slicestring := []string{"A", `D`, "B", `C`, `E`}
	sort.Ints(sliceint)
	fmt.Println(sliceint)
	sort.Strings(slicestring)
	fmt.Println(slicestring)
}