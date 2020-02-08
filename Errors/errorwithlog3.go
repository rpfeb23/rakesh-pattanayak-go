package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	// Below will run as FATAL happened
	defer foo()


	f,err := os.Create("log-errorwithlog3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	nosuchfile , err := os.Open("Nosuchfile.txt")
	if err != nil {
		// for Fatal os.Exit will be called and defer functions will NOT run fatal is same as Println with os.Exit 1 call
		log.Fatalln(" Fatal while opening : ",err)
	}
	defer nosuchfile.Close()

	// Below will run as FATAL happened earlier
	fmt.Println(`Check the file "log-errorwithlog3.txt" for logs`)

}

func foo()  {
	fmt.Println(" This defered function will not run with Defered")
}