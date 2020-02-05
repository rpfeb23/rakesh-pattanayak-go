package DatatypesLoops

import "fmt"

const (
	a = iota + 1
	b = iota
	c = iota
)
// Only when iota has a const keyword infront of it it resets
const (
	x = iota
	y
	z
)

const p = iota
const q = iota

func main()  {
	fmt.Println(a, "\t", b , "\t", c)
	fmt.Printf("%T\t%T\t%T\n",a,b,c)

	fmt.Println(x, "\t", y , "\t", z)
	fmt.Printf("%T\t%T\t%T\n",a,b,c)

	fmt.Println(p, "\t", q )
	fmt.Printf("%T\t%T\n",p,q)
}
