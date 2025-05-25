package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jdkato/prose/v2"
)

func main() {
	// Get the file path from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run typo-check.go <file-path>")
		return
	}

	filePath := os.Args[1]

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a new document from prose
	for scanner.Scan() {
		line := scanner.Text()

		// Tokenize the text into words and check for spelling errors
		doc, err := prose.NewDocument(line)
		if err != nil {
			fmt.Println("Error processing line:", err)
			continue
		}

		// Check for typos in each token (word)
		for _, token := range doc.Tokens() {
			// Simple check: consider a word as a typo if it contains a non-alphabetic character
			if !isCorrectWord(token.Text) {
				fmt.Printf("Typo found: %s\n", token.Text)
			}
		}
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// Check if the word is correct (basic check for non-alphabetic characters)
func isCorrectWord(word string) bool {
	for _, r := range word {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}
