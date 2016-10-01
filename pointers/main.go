package main

import "fmt"

func main() {

	// Declare an integer variable with the value of 20.
  	count := 20
	// Display the address of and value of the variable.
    fmt.Println("Value", count)
    fmt.Println("Address", &count)
	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
      p := &count
	// Display the address of, value of and the value the pointer
	// points to.
    fmt.Println("Address Of:", &p, "Value Of:", p, "Points To:", *p)
}
