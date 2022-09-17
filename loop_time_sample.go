package main

// Import required packages
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// Define main function
func main() {

	// Open file
	f, err := os.Open("data/full_matrix.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	defer f.Close()

	// Read csv into 2-d array
	csvReader := csv.NewReader(f)
	arr, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate loop start time
	var start_time = time.Now()

	// Accessing each value of the array
	fmt.Println("Elements of my array")
	for p := 0; p < len(arr); p++ {
		for q := 0; q < len(arr[0]); q++ {
			fmt.Println(arr[p][q])
		}
	}

	// Calculate loop end time
	var end_time = time.Now()

	// Print time to run
	fmt.Println("Time to run loop:", end_time.Sub(start_time))

}
