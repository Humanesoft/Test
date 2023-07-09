package main

import (
	"fmt"
	"github.com/Humanesoft/Test"
)

func main() {
	total := calculator.Sum(4, 7)
	fmt.Println(total)
	fmt.Println("Version:", calculator.Version)
}
