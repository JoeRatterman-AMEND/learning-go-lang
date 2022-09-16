package main

// Import required packages
import "fmt"

// Define main function
func main() {

	// Define variable age
	var age int

	// Declare age value
	age = 27

	// Print age
	fmt.Println("My age is", age)

	// Decare multiple variables name & age - implied type
	name, new_age := "joe", 27

	// Print new values for name & new_age
	fmt.Println("my name is", name, "age is", new_age)

}
