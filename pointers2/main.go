// Add imports.
package main

import "fmt"

// Declare a type named user.
type user struct {
  name string
  email string
  age int
}

// Create a function that changes the value of one of the user fields.
func age(u *user, age int) {
  	u.age = age
  }

func main() {
	// Create a variable of type user and initialize each field.
  bob := user{"Bob", "bob@bob.com", 25}
	// Display the value of the variable.
  fmt.Println(bob)
	// Share the variable with the function you declared above.
  age(&bob, 26)
	// Display the value of the variable.
  fmt.Println(bob)
}
