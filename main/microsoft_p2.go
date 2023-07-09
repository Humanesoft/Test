package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	Sunday = iota + 1

	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}
func main() {
	// var integer int = 2147483648
	// fmt.Println(integer)
	// fmt.Println("Monday:", Monday)
	// fmt.Println("Tuesday:", Tuesday)
	// fmt.Println("Wednesday:", Wednesday)
	// fmt.Println("Thursday:", Thursday)
	// fmt.Println("Friday:", Friday)
	// fmt.Println("Saturday:", Saturday)
	// fmt.Println("Sunday:", Sunday)
	// i, _ := strconv.Atoi("-42")
	// s := strconv.Itoa(-42)
	// fmt.Println(i, s)
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
}
