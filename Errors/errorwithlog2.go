package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	f,err := os.Create("log-errorwithlog2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	nosuchfile , err := os.Open("Nosuchfile.txt")
	if err != nil {
		log.Println(" Error while opening ",err)
	}
	defer nosuchfile.Close()

	fmt.Println(`Check the file "log-errorwithlog2.txt" for logs`)

}