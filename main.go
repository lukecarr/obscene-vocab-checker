package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// getFilename reads a filename string from stdin
func getFilename() (filename string) {
	fmt.Scan(&filename)
	return filename
}

// iterFileLines opens a file for reading, and iterates over all
// lines, invoking an iterFunc on each line
func iterFileLines(filename string, iterFunc func(line string)) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		iterFunc(scanner.Text())
	}
}

// sanitize converts a word (string) to lowercase and removes
// all whitespace padding
func sanitize(word string) string {
	return strings.TrimSpace(strings.ToLower(word))
}

// getWords returns a set of taboo words from a file
func getWords(filename string) map[string]bool {
	words := make(map[string]bool)

	iterFileLines(filename, func(line string) {
		words[sanitize(line)] = true
	})

	return words
}

// censorWord checks if a word is in the set of
// forbidden/taboo words, and censors with asterisks if so
func censorWord(word string, words map[string]bool) string {
	if !words[sanitize(word)] {
		return word
	}

	return strings.Repeat("*", utf8.RuneCountInString(word))
}

// iterSentence iterates over a sentence's words, ignoring
// punctuation, and performs the iterFunc to each word
func iterSentence(sentence *string, iterFunc func(word string)) {
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	for _, word := range strings.Fields(replacer.Replace(*sentence)) {
		iterFunc(word)
	}
}

func main() {
	filename := getFilename()
	words := getWords(filename)

	var sentence string

	for {
		fmt.Scanln(&sentence)

		if sentence == "exit" {
			break
		}

		iterSentence(&sentence, func(word string) {
			censored := censorWord(word, words)
			sentence = strings.ReplaceAll(sentence, word, censored)
		})

		fmt.Println(sentence)
	}

	fmt.Println("Bye!")
}
