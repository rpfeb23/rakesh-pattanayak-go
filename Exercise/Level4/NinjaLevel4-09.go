package main

import "fmt"

func main()  {
	m := map[string][]string{
		"bond_james":{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss":{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m["Rakesh"] = []string{"Adi","Driving"}
	fmt.Println(m)
}