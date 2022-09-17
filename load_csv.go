package main

// Import required packages
import (
	"fmt"
	"log"
	"os"

	"github.com/tobgu/qframe"
)

// Define main function
func main() {

	// Define csv file path
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Load csv to df
	f := qframe.ReadCSV(csvfile)

	// Print df rows
	fmt.Println(f)

}
