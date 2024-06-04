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
	log.Infof("Words in dictionary %s: %d", *dict, words.Len())
	fragments := strings.ToLower(flag.Args()[0])
	// Split the input into words
	wordsInFragments := strings.Fields(fragments)
	fSet := sets.FromSlice(wordsInFragments)
	// Format like quartiles original screenshot:
	log.Infof("Looking at fragments:")
	for i, word := range wordsInFragments {
		fmt.Printf("%s\t", word)
		if i%4 == 3 {
			fmt.Println()
		}
	}
	// Check all combinations of 2-4 words
	for i := 2; i <= 4; i++ {
		log.Infof("Checking %d words combinations:", i)
		for _, c := range sets.Tuplets(fSet, i) {
			w := strings.Join(c, "")
			if words.Has(w) {
				fmt.Println(w)
			}
		}
	}
}
