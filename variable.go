// Assign values to variables

package main
import "fmt"

func main() {
 
  // explicitly declaring the data type and assigning value
  var number1 int = 10
  fmt.Println(number1)

 // Assigning a value without explicitly declaring the data type
  var number2 int = 20
  fmt.Println(number2)

  // shorthand notation to define variable
  number3 := 30
  fmt.Println(number3)  
  number2= 50; 
  fmt.Println("The next value is : ", number2)
}
