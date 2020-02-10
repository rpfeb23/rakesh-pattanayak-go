package mymath

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type input struct {
		data []int
	}

	inputset := []input{
		{data:[]int{1,2,3,4,5}},
		{data:[]int{100,200,300,400,500}},
	}

    expected := []float64{3,300}

    got := []float64{}

    for _, v := range inputset{
    	got = append(got,CenteredAvg(v.data))
	}

    if (reflect.DeepEqual(expected,got)){
    	fmt.Println(" Test PASSED")
	}else{
		t.Error(" Test Failed ")
	}
	fmt.Println("Expected : ", expected, " Got : ", got)
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i :=0 ; i < b.N; i++{
		CenteredAvg([]int{1,2,3,4,5})
	}
}

func ExampleCenteredAvg() {
	{
		input := []int{}
		input = []int{1, 2, 3, 4, 5}
		fmt.Println(CenteredAvg(input))
	}
	// Output: {
	// 3
	// The smallest value 1 and largest value 5 are taken out
	// now the slice is [2,3,4]
	// whose average is (2+3+4)/count = 3
	// }
}