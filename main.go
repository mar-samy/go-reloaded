package main

import (
	f "fmt" // will be using f as as shortcut instead writing the whole fmt term everytime i need it
	"os"
	"strconv"
	s "strings" // will be using s as a shortcut instead of writing strings everytime i use the functions in that package
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
	err = os.WriteFile(outputFile, []byte(content), 0o622)
	if err != nil {
		f.Println("Error to write the file", err)
		return
	}
}

func processText(content string) string {
	words := s.Fields(content)
	words = hexToDecimal(words)
	words = binToDecimal(words)
	words = upperModifier(words)
	words = lowerModifier(words)
	words = capitalizeModifier(words)
	words = upperModNum(words)
	words = lowerModNum(words)
	words = capModNum(words)
	words = fixPunctuation(words)
	words = fixArticles(words)
	content = s.Join(words, " ")
	return content
}

func hexToDecimal(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			decimalNumber, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid hexadecimal number")
				continue
			} else {
				words[i-1] = strconv.FormatInt(decimalNumber, 10)
			}
			words[i] = ""
		}
	}
	return words
}

func binToDecimal(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			decimalNumber, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				f.Println("Parsing failed: Invalid bindecimal number")
				continue
			} else {
				words[i-1] = strconv.Itoa(int(decimalNumber))
			}
			words[i] = ""
		}
	}
	return words
}

func upperModifier(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = s.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func lowerModifier(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = s.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func capitalizeModifier(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == "(cap)" {
			word := words[i-1]
			words[i-1] = (s.ToUpper(word[0:1]) + s.ToLower(word[1:]))
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}

func upperModNum(words []string) []string { //(up, 6
	for i := 0; i < len(words); i++ {
		if words[i] == "(up," && i+1 < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				f.Println("Error: Invalid Number Of Words To Be Uppered")
				continue
			}
			if Num > i {
				Num = i
				f.Println("Number of words to be uppered is greater than the available words")
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
	return words
}

func lowerModNum(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low," && i+1 < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				f.Println("Error: Number Of Words To Be Lowered")
				continue
			}
			if Num > i {
				Num = i
				f.Println("Number of words to be lowered is greater than the available words")
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
	return words
}

func capModNum(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," && i < len(words) {
			words[i+1] = s.Trim(words[i+1], ")")
			Num, err := strconv.Atoi(words[i+1])
			if err != nil {
				f.Println("Error: Invalid Number of Words to Be Capitalized")
				continue
			}
			if Num > i {
				Num = i
				f.Println("Number of words to be capitalized is greater than the available words")
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
	return words
}

func fixPunctuation(words []string) []string {
	result := []string{}

	for _, word := range words {
		// If word is a punctuation token (. , ! ? : ; ... !! !?):
		//   Attach it to the last element of result (no space between)
		// Otherwise:
		//   Append it normally
	}

	return result
}

func fixQuotes(words []string) []string {
	// Iterate through words
	// Track whether you are currently inside an open quote or not
	// When you find a standalone ':
	//   If no quote is open: mark the NEXT word as starting with '
	//   If a quote is open: attach ' to the END of the previous word, close the quote
	// Return the cleaned result
}

/*func fixPunctuation(content string) string {

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
//As Elton John said:'I am the most well-known homosexual in the world'
/*words := s.Fields(content)
return words
return content
}
*/

func fixArticles(words []string) []string {
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
			}
		}
	}

	return words
}
