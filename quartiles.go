package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"fortio.org/cli"
	"fortio.org/log"
	"fortio.org/sets"
)

// Get words set.
func getWords(dictionaryPath string) sets.Set[string] {
	// Open the dictionary file
	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatalf("Error opening dictionary file: %v", err)
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := sets.New[string]()
	// Read the file word by word
	for scanner.Scan() {
		word := scanner.Text()
		// skip words lower than 5 characters
		if len(word) < 5 {
			continue
		}
		// normalize the word to lowercase
		word = strings.ToLower(word)
		// add to set
		words.Add(word)
	}
	file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading dictionary file: %v", err)
	}
	return words
}

func main() {
	dict := flag.String("dictionary", "/usr/share/dict/words", "Dictionary file to use")
	cli.MinArgs = 1
	cli.Main()
	words := getWords(*dict)
	fmt.Printf("Words in dictionary: %d\n", words.Len())
}
