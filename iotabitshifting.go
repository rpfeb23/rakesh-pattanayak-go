package main

import "fmt"

const (
	_ = iota      // throwaway 0
	kilobytes = 1 << (iota * 10) // iota is now 1; 1 shifted 10
	megabytes = 1 << (iota * 10) // iota is now 2; 1 shifted 20
	gigabytes = 1 << (iota * 10) // iota is now 3; 1 shifted 30
)

func main()  {

	kb := 1024          // 1 shifted 10 places
	mb := 1024 * kb     // 1 shifted 20 places
	gb := 1024 * mb     // 1 shifted 30 places
	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",mb,mb)
	fmt.Printf("%d\t%b\n",gb,gb)

    fmt.Println("****************** Using IOTA *********************")

	fmt.Printf("%d\t\t%b\n",kilobytes,kilobytes)
	fmt.Printf("%d\t\t%b\n",megabytes,megabytes)
	fmt.Printf("%d\t%b\n",gigabytes,gigabytes)


}
