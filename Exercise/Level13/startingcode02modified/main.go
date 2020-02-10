package main

import (
	"fmt"
	"github.com/rpfeb23/rakesh-pattanayak-go/Exercise/Level13/quote"
	"github.com/rpfeb23/rakesh-pattanayak-go/Exercise/Level13/Word"
)

func main() {
	fmt.Println(Word.Count(quote.SunAlso))

	for k, v := range Word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}