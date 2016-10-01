package main

import "fmt"

// named struct
type user struct {
  name string
  email string
  age int
}

func main() {
  bob := user{"Bob", "bob@bob.com", 25}
  fmt.Println(bob)

  // anonymous struct
  sarah := struct {
    name string
		email string
		age int
	}{
    name: "Sarah",
    email: "sarah@sarah.com",
    age: 22,
  }
  fmt.Println(sarah)
}
