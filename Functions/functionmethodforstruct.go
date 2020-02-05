package main

import "fmt"

type person struct {
	name string
	age int
} 
type secretagent struct{
	person
	agentid int
	licensetokill bool
}

func main()  {
	sa1 := secretagent{
		person:        person{"James Bond",35},
		agentid:       007,
		licensetokill: false,
	}

	sa1.speak() // Notice how we are calling Speak

	talk(sa1)   // Notice the difference
}

// func (r receiver) functionname(params) (returns){code}
// This is method

func (s secretagent) speak()  {
	fmt.Println("I can speak. My name is ", s.name, " ID : ",s.agentid)
}

// Regular function without attaching to a type

func talk(s secretagent)  {
	fmt.Println("I can speak. My name is ", s.name, " ID : ",s.agentid)
}