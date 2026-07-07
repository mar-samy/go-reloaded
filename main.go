package main

import (
	"fmt"
	f "fmt" //will be using f as as shortcut instead writing the whole fmt term everytime i need it
	"os"
	"strconv"
	s "strings" //will be using s as a shortcut instead of writing strings everytime i use the functions in that package
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
	content = processText(content)

	outputFile := os.Args[2]
	err = os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		f.Println("Error to write the file", err)
	}
}
func processText(content string) string {
	content = hexToDecimal(content)
	content = binToDecimal(content)
	content = upperModifier(content)
	content = lowerModifier(content)
	content = capitalizeModifier(content)
	content = upperModNum(content)
	content = lowerModNum(content)
	content = capModNum(content)
	content = fixPunctuation(content)
	content = fixArticles(content)
	return content
}

func hexToDecimal(content string) string {
	words := s.Fields(content)
	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			decimalNumber, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid hexadecimal number")
				return content
			} else {
				words[i-1] = strconv.FormatInt(decimalNumber, 10)
			}
			words[i] = ""
		}
	}
	return s.Join(words, " ")
}

func binToDecimal(content string) string {
	words := s.Fields(content)
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			decimalNumber, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid bindecimal number")
				return content
			} else {
				words[i-1] = strconv.Itoa(int(decimalNumber))
			}
			words[i] = ""
		}
	}
	return s.Join(words, " ")
}
func upperModifier(content string) string {
	words := s.Fields(content)
	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = s.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return s.Join(words, " ")
}
func lowerModifier(content string) string {
	words := s.Fields(content)
	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = s.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return s.Join(words, " ")
}
func capitalizeModifier(content string) string {
	words := s.Fields(content)
	for i := 1; i < len(words); i++ {
		if words[i] == "(cap)" {
			word := words[i-1]
			words[i-1] = (s.ToUpper(word[0:1]) + s.ToLower(word[1:]))
			words = append(words[:i], words[i+1:]...)
		}
	}
	return s.Join(words, " ")
}
func upperModNum(content string) string {
	words := s.Fields(content)
	for i := 0; i < len(words); i++ {
		if words[i] == "(up," && i+1 < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error: Invalid Number Of Words To Be Uppered")
				return content
			}
			if Num > i {
				Num = i
				for j := 0; j < Num; j++ {
					words[(i-1)-j] = s.ToUpper(words[(i-1)-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				for j := 0; j < Num; j++ {
					words[(i-1)-j] = s.ToUpper(words[(i-1)-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return s.Join(words, " ")
}
func lowerModNum(content string) string {
	words := s.Fields(content)
	for i := 0; i < len(words); i++ {
		if words[i] == "(low," && i+1 < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error: Number Of Words To Be Lowered")
				return content
			}
			if Num > i {
				Num = i
				for j := 0; j < Num; j++ {
					words[(i-1)-j] = s.ToLower(words[(i-1)-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				for j := 0; j < Num; j++ {
					words[(i-1)-j] = s.ToLower(words[(i-1)-j])
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return s.Join(words, " ")
}
func capModNum(content string) string {
	words := s.Fields(content)
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," && i < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Println("Error: Invalid Number of Words to Be Capitalized")
				return content
			}
			if Num > i {
				Num = i
				for j := 0; j < Num; j++ {
					word := words[i-1-j]
					words[(i-1)-j] = (s.ToUpper(word[0:1]) + s.ToLower(word[1:]))
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				for j := 0; j < Num; j++ {
					word := words[i-1-j]
					words[(i-1)-j] = (s.ToUpper(word[0:1]) + s.ToLower(word[1:]))
				}
				words = append(words[:i], words[i+2:]...)
			}
		}
	}
	return s.Join(words, " ")
}
func fixPunctuation(content string) string {

	content = s.ReplaceAll(content, " .", ".")
	content = s.ReplaceAll(content, " ,", ",")
	content = s.ReplaceAll(content, " !", "!")
	content = s.ReplaceAll(content, " ?", "?")
	content = s.ReplaceAll(content, " ;", ";")
	content = s.ReplaceAll(content, " : ", ": ")
	content = s.ReplaceAll(content, ",", ", ")
	content = s.ReplaceAll(content, ".", ". ")
	content = s.ReplaceAll(content, "!", "! ")
	content = s.ReplaceAll(content, "?", "? ")
	content = s.ReplaceAll(content, ";", "; ")
	content = s.ReplaceAll(content, ". . .", "...")
	content = s.ReplaceAll(content, "! !", "!!")
	content = s.ReplaceAll(content, "! ?", "!?")
	content = s.ReplaceAll(content, "? !", "?!")
	content = s.ReplaceAll(content, ". '", ".'")
	content = s.ReplaceAll(content, "' ", "'")
	content = s.ReplaceAll(content, " '", "'")
	words := s.Fields(content)
	return s.Join(words, " ")

}
func fixArticles(content string) string {
	words := s.Fields(content)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := words[i+1]
			if nextWord == "" {
				continue
			}
			firstChar := nextWord[0]
			switch firstChar {
			case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			default:
				return words[i]
			}
		}
	}

	return s.Join(words, " ")
}
