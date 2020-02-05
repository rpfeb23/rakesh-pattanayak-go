package main

import "fmt"

type aliens []struct{
	code int
	galaxy string
}

func main()  {
	fmt.Println(aliens{})
	p1 := aliens{
		{
			code:001,
			galaxy:"A",
		},
		{
			code:001,
			galaxy:"B",
		},
	}
	fmt.Println(p1)
    for i:= 0; i<len(p1); i++{
        fmt.Println(p1[i])
        fmt.Printf("%T",p1[i])
	}
    fmt.Println()
    p2 := aliens{{},{}}
    fmt.Println(p2)
	for i:=0;i<2 ;i++  {
		fmt.Println(p2[i])
		p2[i].galaxy = "X"
		p2[i].code = i
	}
    fmt.Println(p2)
}
