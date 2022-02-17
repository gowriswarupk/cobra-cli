package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Println("Enter first value: ")

	fmt.Scan(&a)

	fmt.Println("Enter second value: ")

	fmt.Scan(&b)

	c := a + b

	fmt.Println("sum of values entered is: ", c)
}
