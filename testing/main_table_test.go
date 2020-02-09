package main

import (
	"fmt"
	"testing"
)

func TestTabletestmysum(t *testing.T){
	type test struct {
		input    []int
		expected int
	}
	
	tests := []test{
		test{
			input:    []int{2,3,4},
			expected: 9,
		},
		test{
			input:    []int{4,-5,6},
			expected: 5,
		},
		test{
			input:    []int{ 0,0,0},
			expected: 0,
		},
	}

	for i, v := range tests{

		got := mysum(v.input...)

		if got != v.expected {
			t.Error(" Expected : ", v.expected, " Got : ", got)
		}else{
			fmt.Printf(" Test Case # % v PASSED \n\t\t Test Data input %v \t Test Data Output %v\n", i+1, v.input, v.expected)
		}
	}


}