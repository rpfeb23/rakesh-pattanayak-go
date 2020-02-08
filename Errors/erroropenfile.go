package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	f , err := os.Open("file1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(bs)
		fmt.Println(string(bs))
	}

}