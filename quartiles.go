package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"fortio.org/cli"
	"fortio.org/log"
	"fortio.org/sets"
)

//go:embed words
var dictWords string

// Get words set.
func getWords(dictionaryPath string) sets.Set[string] {
	// Open the dictionary file if provided
	if dictionaryPath == "" {
		// Use the embedded dictionary
		return sets.FromSlice(strings.Fields(dictWords))
	}
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

func hintReplace(word string, idx int) string {
	return word[0:idx] + strings.Repeat("_", len(word)-idx)
}

func main() {
	dict := flag.String("dict", "", "Dictionary file `path` to use, instead of default embedded one")
	foursOnly := flag.Bool("4", false, "Only find the 4-part words")
	startsOnly := flag.Bool("hints", false, "Only show the first segment of the answers (hints only mode)")
	cli.MinArgs = 1
	cli.MaxArgs = -1
	cli.ArgsHelp = "word fragments...\n" +
		"Finds all words in dictionary that can be made from the given fragments"
	cli.Main()
	words := getWords(*dict)
	log.Infof("Words in dictionary %s: %d", *dict, words.Len())
	// Either 1 argument with spaces between fragments or multiple arguments
	fragments := strings.ToLower(strings.Join(flag.Args(), " "))
	// Split the input into words
	wordsInFragments := strings.Fields(fragments)
	fSet := sets.FromSlice(wordsInFragments)
	// Format like quartiles original screenshot:
	log.Infof("Looking at fragments:")
	lastIdx := len(wordsInFragments) - 1
	for i, word := range wordsInFragments {
		fmt.Printf("%s\t", word)
		if i%4 == 3 || i == lastIdx {
			fmt.Println()
		}
	}
	// Check all combinations of 2-4 words (or 4 only if -4 flag)
	start := 2
	if *foursOnly {
		start = 4
	}
	for i := start; i <= 4; i++ {
		log.Infof("Checking %d parts combinations:", i)
		for _, c := range sets.Tuplets(fSet, i) {
			w := strings.Join(c, "")
			if words.Has(w) {
				if *startsOnly {
					fmt.Printf("%s%s\t: %s\n", c[0], strings.Repeat("-*", len(c)-1), hintReplace(w, len(c[0])))
				} else {
					fmt.Printf("%s\t: %s\n", strings.Join(c, "-"), w)
				}
			}
		}
	}
}
