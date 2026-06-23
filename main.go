// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Check that exactly 2 arguments were provided (input and output file)
	if len(os.Args) != 3 {
		fmt.Println("Error: Incorrect number of arguments")
		return
	}
	inputFile := os.Args[1]
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading this file")
		return
	}



	 
	outputFile := os.Args[2]
	content := input
	er := os.WriteFile(outputFile, content, 0644)
	if er != nil {
		fmt.Println("Error to write the file")
	}
	// 2. Read the input file into a string

	// 3. (Transformations will go here in later milestones)

	// 4. Write the result string to the output file
}
