// !!!!! Important !!!!!
// Your Struct Variable names have to start with UPPERCASE inorder for JSON Marshalling to Work

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string              // you cant have "name" First Letter UpperCase
	Age int                  // 1st Letter Uppercase
	Placesvisited []string   // 1st Letter Uppercase
}

func main()  {
	var people []person
    p1 := person{
		Name:          "Rakesh ",
		Age:           35,
		Placesvisited: []string{"India", "USA"},
	}
	p2 := person{
		Name:          "Rajesh ",
		Age:           38,
		Placesvisited: []string{"India", "AUSTRALIA"},
	}
	people = append(people,p1,p2)
	// ANOTHER WAY
	// people := []person{p1,p2,}

	fmt.Println("PEOPLE : ", people)

	// Another Way to populate slice of structs
    var person1 = []person{
		{
			Name:          "Sukanya",
			Age:           32,
			Placesvisited: []string{"XYZ","DEF"},
		},
		{
			Name:          "Sulagna",
			Age:           35,
			Placesvisited: []string{"ABC","PQR"},
		},
	}
    fmt.Println("PERSON : ", person1)

    // JSON Marshallinng Returns :  []byte, error
    // []byte has type []uint8
    json1bytesslice, err := json.Marshal(people)
    if err != nil{
    	fmt.Println(err)
	}
    fmt.Printf("Type of json1bytesslice : %T\n", json1bytesslice)
    fmt.Println("JSON output after converting []byte to string :")
    jsonoutput := string(json1bytesslice)
    fmt.Println(jsonoutput)

}