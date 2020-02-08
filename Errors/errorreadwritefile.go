package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	f, err := os.Create("file1.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("Type of f : %T  Value of f : %v\n", f,f)
	}

	r := strings.NewReader(" James Bond ")
	io.Copy(f,r)
	r = strings.NewReader("\n")
	io.Copy(f,r)
	r = strings.NewReader(" Miss Monney Penny")
	io.Copy(f,r)

}