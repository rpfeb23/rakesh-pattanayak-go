package main

import "fmt"

func main() {

	var s string = `ABCD "Hello,
World"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var bytestring []byte
	bytestring = []byte(s)
	fmt.Println(bytestring)  // Prints the Asci Values
	fmt.Printf("%T\n", bytestring)

	// Print the UTF-8
	for i , v := range  s {
		fmt.Printf("At index %v the value in UTF-8 is %#U \n", i,v)
	}

	for i1, v1 := range bytestring{
		fmt.Printf("At index %v the value bytestring is %v in UTF-8 it is %#U\n", i1,v1,v1)
	}

	//indexing a string accesses individual bytes, not characters.
	// When we iterate over String type it went till Index 10
	// When we iterate over []byte index went till 12 as last two are double bytes

	s1 := "Hello, 世界"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	for i , v := range  s1 {
		fmt.Printf("At index %v the value in UTF-8 is %#U \n", i,v)
	}

	bytestring1 := []byte(s1)
	for i1, v1 := range bytestring1{
		fmt.Printf("At index %v the value bytestring is %v in UTF-8 it is %#U\n", i1,v1,v1)
	}

}
