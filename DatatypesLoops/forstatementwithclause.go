package DatatypesLoops

import "fmt"

func main(){
	for x:=0 ; x < 10; x++{
		fmt.Println(x)
	}

	a := 0
	for ; a <5 ; a++  {
		fmt.Println(a)
	}

	b := 0
	for{
		if b > 3 {
			break
		}
		fmt.Println(b)
		b++
	}
}
