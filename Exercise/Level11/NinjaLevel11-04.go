package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	sqrt1 := sqrtError{
		lat:  "50.2289 N",
		long:  "99.4656 W",
		err:  fmt.Errorf("-Ve Number %v Square Root ", f),
		//This will also work
		//err : errors.New("-ve Number Square Root"),
	}
	if f < 0 {
		return 0, sqrt1
	}
	return 42, nil
}
