package main

import "fmt"

func main() {

// Declare string array with elements initalized to zero value
	var vegetable [5]string

// Declare string array with elements initalized with some values
  fruit := [5]string{"apple", "pear", "orange", "grape", "banana"}

//Assign one array to the other
	vegetable = fruit

//Iterature through array and display string value and address
  for i, item :=range vegetable {
    fmt.Println(item, &vegetable[i])
  }

}
