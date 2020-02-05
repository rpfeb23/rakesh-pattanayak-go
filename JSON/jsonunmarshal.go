// Input JSON is from jsonmarshal
//[{"Name":"Rakesh ","Age":35,"Placesvisited":["India","USA"]},{"Name":"Rajesh ","Age":38,"Placesvisited":["India","AUSTRALIA"]}]

// use https://mholt.github.io/json-to-go/ to get your struct from JSON

package main

import (
	"encoding/json"
	"fmt"
)

type people []struct {
	Name          string   `json:"Name"`
	Age           int      `json:"Age"`
	//json name and struct name different below (1st Letter Uppercase)
	Places        []string `json:"Placesvisited"`
}

func main()  {
	// capture your json in a string
    inputjson := `[{"Name":"Rakesh ","Age":35,"Placesvisited":["India","USA"]},{"Name":"Rajesh ","Age":38,"Placesvisited":["India","AUSTRALIA"]}]`

    p1 := people{}
    fmt.Println("Value of P1 before Unmarshalling : ", p1)
    err := json.Unmarshal([]byte(inputjson),&p1)
    if err != nil{
    	fmt.Println(err)
	}
    fmt.Println("Value of P1 after Unmashalling : ", p1)

    for _, v := range p1{
    	fmt.Println(v)
    	fmt.Println("\t\t ",  v.Name)
    	fmt.Println("\t\t  ", v.Age)
    	fmt.Println("\t\t  ", v.Places)
	}
}