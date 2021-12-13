package main

import (
	"fmt"
)

func main() {
	// Data types in go are
	//  int , float , strings , bool , byte
	// complex (2+4i)
	// rune 	Used for characters. Internally used as 32-bit integers. 'a', '7', '<'
	var integer1 int
	var integer2 int
	integer1 = 5
	integer2 = 10
    currentAge := 20
	fmt.Println(integer1)
	fmt.Println(integer2)
	fmt.Printf("Age = %d\n", currentAge)
	// # Program to illustrate float32 and float64 with example
	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901

	// can store decimals with greater precision
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)
	var message string
	message = "Welcome to Nit patna"

	fmt.Println(message)
	var boolValue bool
	boolValue = false

	fmt.Println(boolValue)
}
