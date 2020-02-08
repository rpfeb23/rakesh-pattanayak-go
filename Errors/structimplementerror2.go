package main

import (
	"fmt"
)

type person struct {
	errorCode int
	errorDesc string
}

func (p *person) Error() string  {
	return fmt.Sprintf(" Error Occured : Code  :%v Desc : %v", p.errorCode, p.errorDesc)
	
}

func main()  {

	p1 := person{
		errorDesc: "_ve Number cant determine sqrt",
		errorCode:  001,
	}

	f, err := squareroot(-10,p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}

func squareroot(f float64, p person) (float64, error) {
	if (f < 0) {
		return 0, &p  // Observe if you will give just P it will complain saying Erro () has a pointer receiver
	}
	return 0, nil
}