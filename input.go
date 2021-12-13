package main

import (
	"fmt"
)
func main(){
	  var collgeName string 
	  roll := 0
	 fmt.Print("Enter the name of your college : ")
	 fmt.Scanln(&collgeName); 
     fmt.Println("Enter the college roll number : ")
	 fmt.Scanln(&roll) 
	 fmt.Println("your roll : ", roll)
	 fmt.Println("Your college name is  : ",collgeName)
}