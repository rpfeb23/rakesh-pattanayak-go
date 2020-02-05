package main

import "fmt"

func main()  {

	var multidimensionslice [][]string

	multidimensionslice = [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Println(multidimensionslice)

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}



	var x =[][]int{}
	fmt.Println(x)
    x = make([][]int,5)
	fmt.Println(x)
    fmt.Println("*********POPULATE MULTIDIMENSIONAL SLICE***********")
    for i,v := range x{
    	fmt.Print(i,v)
    	z := i + 3
    	x[i] = make([]int,z)
    	for j,_ := range x[i]{
    		x[i][j] = i + j
    		fmt.Printf("\t\t%v",x[i][j])
		}
    	fmt.Println()
	}
    fmt.Println("How Entire Slice Looks ", x, "Length : ", len(x))
    fmt.Println("******** ITERATE OVER Multi Dimensional Slice - OPtion1 **************")
	for index1:=0;index1<len(x);index1++{
		fmt.Print("Outer Value : ",x[index1])
		for index2:=0;index2<len(x[index1]);index2++{
			fmt.Print(" Inner Value : ",x[index1][index2])
		}
		fmt.Println()
	}

	fmt.Println("******** ITERATE OVER Multi Dimensional Slice - OPtion2 **************")

	for _, v := range x{
		fmt.Print("Outer Value : ",v)
		for _,innervalue := range v{
			fmt.Print("  Inner Value : ", innervalue)
		}
        fmt.Println()
	}
}
