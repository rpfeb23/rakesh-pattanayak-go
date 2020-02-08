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

func (p person) Error() string  {
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
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	err = fmt.Errorf("Fudged", a)
	fmt.Println(err)
	fmt.Println("--------------------------------------------")
	if err != nil {
		err = fmt.Errorf(a.(person).Error(),a)
	}
	return bs,err
}
