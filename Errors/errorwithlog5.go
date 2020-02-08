package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	defer foo()
	f,err := os.Create("log-errorwithlog5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	nosuchfile , err := os.Open("Nosuchfile.txt")
	if err != nil {
		// panic function will call Defered functions but not write to log file
		panic(" Fatal while opening : ")
	}
	defer nosuchfile.Close()

	// Below will run as PANIC happened earlier

	fmt.Println(`Check the file "log-errorwithlog5.txt" as we use panic nothing should be written`)

}

func foo()  {
	fmt.Println(" This defered function will run as we using PANIC")
}