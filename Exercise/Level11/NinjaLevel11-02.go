package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func (p *person) Error() string  {
	customerror := fmt.Sprintf(" Error Occured While Marshalling : Firs %v : Last : %v : Sayings : %v :", p.First, p.Last,p.Sayings)
	return customerror
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println(string(bs))
	}

}

// toJSON needs to return an error also
func toJSON(p person) ([]byte, error) {
	bs, err := json.Marshal(p)
	fmt.Println(" Standard Error :", err)
	if err != nil {
		err = fmt.Errorf( p.Error() , p )
	}
	return bs,err
}
