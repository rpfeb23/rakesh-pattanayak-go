package main

import "fmt"

func main() {
	m := map[string]int{
		"Rakesh": 85050,
		"Rajesh": 754005,
	}

	fmt.Println(m)
	fmt.Println(m["Rakesh"])

	// Even if Rojalin is not there it prints Zero Value
	fmt.Println(m["Rojalin"])

	// Handle with , ok Option 1 IF Style
	v, ok := m["Rojalin"]
	if ok {
		fmt.Println("Rojalin Present in Map : ", v)
	} else {
		fmt.Println("Rojalin NOT Present in Map : ")
	}
	// Handle with , ok Option 1 IF Style
	if v1, ok1 := m["Rakesh"]; ok1 {
		fmt.Println("Rakesh Present in Map : ", v1)
	}

	// Map with SLICES in Value

	m1 := map[string][]int{
		"Rakesh": {85050, 860},
		"Rajesh": {754005, 98610},
	}
	fmt.Println(m1)
	fmt.Println(m1["Rakesh"])

	for i, v := range m1 {
		fmt.Println("Index: ", i, "Value :", v)
	}

	// map with another map in value
	m2 := map[string]map[string]int{
		"Rakesh": {"USA": 85050,
			       "India": 98610},
		"Sukanya": {"USA": 85050,
			        "India": 94373},
	}
	fmt.Println(m2)
	for i, v := range m2{
		fmt.Print(i,v)
		fmt.Print("      ", v["India"], "      ")
		for i1, v1:= range v{
			fmt.Print(i1, v1)
		}
		fmt.Println()
	}
}
