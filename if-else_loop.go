package main

// Import required packages
import "fmt"

// Define main function
func main() {

	// Define 4x3 array
	arr := [4][3]int{
		{100, 5, 9},
		{95, 23, 43},
		{23, 94, 44},
		{84, 234, 934}}

	// Print full array
	fmt.Println("Below are the elements in my array:")
	fmt.Println(arr)

	// Loop to view previous rows' data
	fmt.Println("Rows of my array")
	for p := 0; p < len(arr); p++ {

		if p == 0 {
			fmt.Println("1st Row!")
		} else {
			fmt.Println("The previous row values are below: ")
			fmt.Println(arr[p-1])
		}

	}

	// Loop to view next rows' data
	fmt.Println("Rows of my array")
	for p := 0; p < len(arr); p++ {

		if p == len(arr)-1 {
			fmt.Println("Last Row!")
		} else {
			fmt.Println("The next row values are below: ")
			fmt.Println(arr[p+1])
		}

	}

}
