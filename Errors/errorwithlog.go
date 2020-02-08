package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	_ , err := os.Open("Nosuchfile.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
}