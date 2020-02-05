package DatatypesLoops

import "fmt"

func main()  {

	// break example
	b := 0
	for {
		if b > 3{
			break
		}
		fmt.Println(b)
		b++
	}


	// Print even number till 10 to emonstrate break and continue
	fmt.Println("****************************************************")
	a := 0
	for {
        a++               // Key is to increment before CONTINUE
		if a > 10 {
			break         // breaks out of loop
		}
		if (a % 2 == 0){
			fmt.Println(a)
			continue      // goes to next iteration of loop
			              // skipping anything below it
			fmt.Println("This line will never be printed")
		}

		//a++   //DANGER for even numbers this statement wil NOT execute
		      // and Might go for INFINITE Loop
	}


}
