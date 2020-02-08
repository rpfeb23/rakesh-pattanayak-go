// trying go run, go build , go install

// go build NinjaLevel9-06.go created one executable and put it in same folder with name NinjaLevel9-06. To execute that file "./NinjaLevel9-06"

// go install will take ur folder and create one executable and PLACE it uner GOPATH/bin

package main
import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}