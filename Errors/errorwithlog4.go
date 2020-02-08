package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	defer foo()
	f,err := os.Create("log-errorwithlog4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	nosuchfile , err := os.Open("Nosuchfile.txt")
	if err != nil {
		// Panic will call Defered functions
		log.Panic(" Fatal while opening : ",err)
	}
	defer nosuchfile.Close()

	// Below will run as PANIC happened earlier

	fmt.Println(`Check the file "log-errorwithlog4.txt" for logs`)

}

func foo()  {
	fmt.Println(" This defered function will run as we using PANIC")
}