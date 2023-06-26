package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func getWordsArray(s string) []string {
	specificCharacters := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	spaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	formattedString := specificCharacters.ReplaceAllString(s, " ")
	formattedString = spaces.ReplaceAllString(formattedString, " ")
	trimString := strings.Trim(formattedString, " ")
	wordsArray := strings.Split(trimString, " ")
	return wordsArray
}

func findFirstUniqueLetterInWord(word string) string {
	var letter string
	for i := 0; i < len(word); i++ {
		var counter uint
		for j := i; j < len(word); j++ {
			if word[i] == word[j] {
				counter++
			}
		}
		if counter == 1 {
			letter = string(word[i])
			break
		}
	}
	return letter
}

func findFirstUniqueLetterInWords(words []string) string {
	lettersInWords := make([]string, 0)
	lettersMap := make(map[string]uint)

	for _, word := range words {
		l := findFirstUniqueLetterInWord(word)
		lettersMap[l]++
		lettersInWords = append(lettersInWords, l)
	}

	for _, letter := range lettersInWords {
		if lettersMap[letter] == 1 {
			return letter
		}
	}
	return "Could not find unique letter in words"
}

func readFile(filename string) (string, error) {
	var text string
	file, err := os.Open(filename)
	if err != nil {
		return text, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}
	return text, nil
}

func main() {
	text, err := readFile(os.Args[1])
	if err != nil {
		log.Fatal("Something wrong with file: ", err)
	}
	if len(text) == 0 {
		log.Fatal("File is empty")
	}
	wordsArray := getWordsArray(text)
	fmt.Println("First unique letter:", findFirstUniqueLetterInWords(wordsArray))
}
