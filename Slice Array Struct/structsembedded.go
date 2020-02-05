package main

import "fmt"

type human struct {
	name string
	age int
}

type ussecrentagent struct{
	secretagentperson human
	agent_id int
}

type uksecrentagent struct{
	human  //We havn't given fieldname.Unqualified typename acts:field name
	agent_id int
}

func main()  {
	p1 := human{
		name: "James Bond",
		age:  37,
	}

	p2 := human{
		name: "Rakesh",
		age:  35,
	}

	sa1 := ussecrentagent{
		secretagentperson: p1,
		agent_id:          007,
	}

	sa2 := uksecrentagent{
		human : human{"Jason Bourne", 32},
		agent_id:          001,
	}

	fmt.Println(p1, p1.age, p1.name)
	fmt.Println(p2)
	fmt.Println(sa1,
		sa1.secretagentperson,
		sa1.secretagentperson.name,
		sa1.secretagentperson.age,
		sa1.agent_id)

	fmt.Println(sa2,
		sa2.human,
		sa2.name,    // You can now directly access the underying struct
		sa2.human.age, // Unqualified typename acts:field name
		sa2.agent_id)
}
