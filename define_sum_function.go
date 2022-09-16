package main

// Import required packages
import "fmt"

// Define main function
func main() {

	// Define array to sum
	arr := []int{100, 1000, 343, 123}

	// Calculate sum
	fmt.Println("My array sum = ", addArray(arr...))

}

// Define sum function
func addArray(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
