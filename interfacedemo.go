// polymerphism
// Assertion
// EVERY TYPE implements EMPTY INTERFACE interface{}
//             e.g. int is also type interface{}
//                  bool is also type interface{}
package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}
func (p person) speak()  {
	fmt.Println("I can Also speak as person struct. My name is ", p.name, " age : ",p.age)
}
type secretagent struct{
	person
	agentid int
	licensetokill bool
}
// func (r receiver) functionname(params) (returns){code}
// This is method

func (s secretagent) speak()  {
	fmt.Println("I can speak. My name is ", s.name, " ID : ",s.agentid)
}

type human interface {
	speak()
}

func bar(h human)  {
	switch  h.(type){
	case person:
		fmt.Println("Type Person is passed as Interface")
		fmt.Println(h.(person).name,h.(person).age)
	case secretagent:
		fmt.Println("Type Person is passed as Interface")
		fmt.Println(h.(secretagent).name,
			        h.(secretagent).age,
			        h.(secretagent).agentid,
			        h.(secretagent).licensetokill)
	}
	fmt.Println("I am human", h)
}


func main()  {
	sa1 := secretagent{
		person:        person{"James Bond",35},
		agentid:       007,
		licensetokill: false,
	}

	sa1.speak()

	// A Value can be of more than One Type
	// sa1 (Value) is of
	//                   1) type secretagent
	//                   2) type human

    p1 := person{
		name: "Jason Bourne",
		age:  32,
	}
    p1.speak()
    fmt.Println("*******************************************")
    fmt.Println("person callling bar passing person param : ")
    bar(p1)
    fmt.Println()
	fmt.Println("Secret Agent callling bar passing SecretAgent param : ")
    bar(sa1)
}
