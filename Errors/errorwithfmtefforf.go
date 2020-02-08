package main

import (
	"fmt"
)

func main()  {

	r, err := squareroot(-10)
	fmt.Println(r,err)
}

func squareroot(f float64) (float64, error) {
	if (f < 0){
		return 0, fmt.Errorf("Negative Number %v : Square root not possible", f)
	}
	return 0, nil
}