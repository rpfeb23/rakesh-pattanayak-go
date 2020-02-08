package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")

	fmt.Printf("%T : \n", err)

}
