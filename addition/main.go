package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Println("Enter first value: ")

	fmt.Scan(&a)

	fmt.Println("Enter second value: ")

	fmt.Scan(&b)

	c = add_values(a, b)

	fmt.Printf("sum of values entered is: ",
		c) //needs updating after rechecking
}

func add_values(a int, b int) int {
	var c = a + b
	return c
}
