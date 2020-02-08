package main

import (
	"errors"
	"fmt"
)
//var error1 = errors.New("Square toot of -ve number not possible")
func main()  {

	r, err := squareroot(-10)
	fmt.Println(r,err)
}

func squareroot(f float64) (float64, error) {
	if (f < 0){
		return 0, errors.New("Negative Number Square root not possible")
		// return 0, error1
	}
	return 0, nil
}