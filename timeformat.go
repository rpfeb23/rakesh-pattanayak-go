package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("Mon"))
	fmt.Println(time.Now().Format("Jan"))
	fmt.Println(time.Now().Format("02"))
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(time.Now().Format("-0700"))
	fmt.Println(time.Now().Format("MST"))
	fmt.Println(time.Now().Format("2006"))
	fmt.Println(time.Now().Format("2006 Mon 02 Jan MST"))
	fmt.Println()
	fmt.Println()
	// There are some predefined constants
	fmt.Println(time.Now().Format(time.Kitchen))
	fmt.Println(time.Now().Format(time.RFC3339))
}