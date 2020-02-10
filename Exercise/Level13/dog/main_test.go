package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	expected := 28
	got := Years(4)

	if got != expected{
		t.Error(" Expected : % v", expected, " Got : %v", got)
	}

}

func ExampleYears() {
	{
		dogyears := Years(4)
		fmt.Println(dogyears)
	}
	//Output:
	//28
}

func BenchmarkYears(b *testing.B) {
	for i:=0;i< b.N ; i++ {
		Years(5)
	}
}

func TestYearsTwo(t *testing.T) {
	expected := 28
	got := Years(4)

	if got != expected{
		t.Error(" Expected : % v", expected, " Got : %v", got)
	}

}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))

	//Output:
	//28
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0;i< b.N ; i++ {
		YearsTwo(5)
	}
}