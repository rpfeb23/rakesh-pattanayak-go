package Word

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseCount(t *testing.T) {
	inputstring := " I am Rakesh. Rakesh is awesome. I am doing great"
	expected := make(map[string]int)
	expected = map[string]int{
		"I" : 2,
		"am" : 2,
		"Rakesh" : 1,
		"Rakesh." : 1,
		"is" : 1,
		"awesome.": 1,
		"doing" : 1,
		"great" : 1,
	}

	got := UseCount(inputstring)

	if (reflect.DeepEqual(expected,got)){
		fmt.Println(" Testing Pass for UseCount")
	}else{
		t.Error(" FAILED Test UseCount :  ", "\n\t\t  Expected : ", expected, " Got : ", got)

	}

}

func BenchmarkUseCount(b *testing.B) {
	for i:=0; i < b.N; i++{
		UseCount("I am Rakesh")
	}
}

func TestCount(t *testing.T) {
	inputstring := " I am Rakesh."
	expected := 13
	got := Count(inputstring)
	if expected != got{
		t.Error("Count Test Failed --> Expected : ", expected, " Got :", got )
	}else{
		fmt.Println(" Testing Pass for Count")
	}
}

func ExampleCount() {
	fmt.Println(Count("I am Rakesh"))
	//Output:
	//11
}

func BenchmarkCount(b *testing.B) {
	for i:=0; i< b.N;i++{
		Count("I am Rakesh")
	}
}