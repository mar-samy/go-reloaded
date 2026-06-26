package main

import (
	f "fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {

	if len(os.Args) != 3 {
		f.Println("Error: Incorrect number of arguments")
		return
	}
	inputFile := os.Args[1]
	input, err := os.ReadFile(inputFile)
	if err != nil {
		f.Println("Error reading this file", err)
		return
	}
	content := string(input)

	content = hexToDecimal(content)
	content = binToDecimal(conetnt)
	/*n, err := strconv.ParseInt("GG", 16, 64)
	fmt.Println(n)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Parsing failed")
		return
	}
	*/
	outputFile := os.Args[2]

	err = os.WriteFile(outputFile, input, 0644)
	if err != nil {
		f.Println("Error to write the file", err)
	}

}
func hexToDecimal(content string) string {
	// Parse s as a base-16 integer
	words := s.Fields(content)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {

			decimalNumber, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid hexadecimal number")
				return content
			} else {
				words[i-1] = strconv.Itoa(int(decimalNumber))
			}
			words[i] = ""
		}
	}
	return s.Join(words, " ")
}

func binToDecimal(content string) string {
	// Parse s as a base-2 integer
	words := s.Fields(content)

	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			decimalNumber, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid hexadecimal number")
				return content
			} else {
				words[i-1] = strconv.Itoa(int(decimalNumber))
			}
			words[i] = ""
		}
		return s.Join(words, " ")
	}

	// If parsing fails, return s unchanged
	// Convert the result back to a decimal string and return it
}
