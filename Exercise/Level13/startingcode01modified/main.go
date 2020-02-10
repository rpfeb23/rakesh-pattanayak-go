package starting_code

import (
	"fmt"
	"github.com/rpfeb23/rakesh-pattanayak-go/Exercise/Level13/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}