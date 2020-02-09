package greeting

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")

	if s != " Hello Dear : James"{
		t.Error(`" Expected  Hello Dear : Rakesh"`)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	//Output:
	//  Hello Dear : James
}

func BenchmarkGreet(b *testing.B) {
	for i :=0 ; i < b.N ; i++{
		Greet("James")
	}
}
