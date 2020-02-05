package DatatypesLoops

import "fmt"

func main() {

	switch {  // nothing in expr after switch means evaluate against TRUE
	case false:
		fmt.Println("This wont print")
	case (2 == 3):
		fmt.Println("Two is equal to three. Wow")
	case (3 == 3):
		fmt.Println("Three is equal to Three")
	case (4 == 4):
		fmt.Println("This will NOT Print even if true")
	}

	fmt.Println("*************** SWITCH FALLTHROUGH **************")
	// switch with fallthrough

	switch {
	case false:
		fmt.Println("This wont print")
	case (2 == 3):
		fmt.Println("Two is equal to three. Wow")
	case (3 == 3):
		fmt.Println("Three is equal to Three")
		fallthrough

	case (4 == 4):
		fmt.Println("This will NOW Print even if previous case is True")
	case (5 == 6):
		fmt.Println("Five is equal to Six. WOW")

	}
	fmt.Println("*************** DEFAULT WONT RUN **************")
	switch {
	case false:
		fmt.Println("This wont print")
	case (2 == 3):
		fmt.Println("Two is equal to three. Wow")
	case (3 == 4):
		fmt.Println("This will Print")
	case (4 == 4):
		fmt.Println("Four is equal to Four")
	default: // prints only if none of cases are TRUE
		fmt.Println("This is DEFAULT. But Won't Print")
	}
	fmt.Println("*************** DEFAULT WILL RUN **************")
	switch {
	case false:
		fmt.Println("This wont print")
	case (2 == 3):
		fmt.Println("Two is equal to three. Wow")
	case (3 == 4):
		fmt.Println("This will Print")
	case (4 == 5):
		fmt.Println("Four is equal to Four")
	default: // prints only if none of cases are TRUE
		fmt.Println("None of cases are TRUE. Printing Default")
	}

	fmt.Println("*************** DEFAULT WITH FALLTHROUGH **************")
	switch {
	case false:
		fmt.Println("This wont print")
	case (2 == 3):
		fmt.Println("Two is equal to three. Wow")
	case (3 == 4):
		fmt.Println("This will Print")
	case (4 == 4):
		fmt.Println("Four is equal to Four")
		fallthrough
	default: // prints only if none of cases are TRUE
		fmt.Println("Printing Default even if there is a TRUE case due to FALLTHROUGH")
	}

	fmt.Println("*************** SWITCH WITH EXPRESSION **************")

	var  day_of_week string = "friday"
	switch day_of_week {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Friday", "friday":
		fmt.Println("Today is Friday(friday)")
	}

}
