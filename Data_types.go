package main
import ("fmt")
func main() {
	// Data types in go are
	//  int , float , strings , bool , byte
	// complex (2+4i)
	// rune 	Used for characters. Internally used as 32-bit integers. 'a', '7', '<'
	var integer1 int
	var integer2 int
	integer1 = 5
	integer2 = 10

	fmt.Println(integer1)
	fmt.Print(integer2)
	// # Program to illustrate float32 and float64 with example
	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901

	// can store decimals with greater precision
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)

}
