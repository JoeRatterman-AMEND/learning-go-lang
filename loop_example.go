package main

// Import required packages
import "fmt"

// Define main function
func main() {

	// Define 3x3 array
	arr := [4][3]int{
		{100, 5, 9}, 
		{95, 23, 43},
        {23, 94, 44}, 
		{84, 234, 934}}

	// Print full array
	fmt.Println("Below are the elements in my array:")
	fmt.Println(arr)

	// Accessing each row of the array
    fmt.Println("Rows of my array")
    for p := 0; p < len(arr); p++ {
        fmt.Println(arr[p])
    }

	// Accessing each value of the array
    fmt.Println("Elements of my array")
    for p := 0; p < len(arr); p++ {
        for q := 0; q < len(arr[0]); q++ {
            fmt.Println(arr[p][q])
        }
    }

}