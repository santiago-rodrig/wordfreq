package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()

		if unicode.IsPunct(rune(word[len(word)-1])) {
			word = word[:len(word)-1]
		}

		words[word]++
	}

	for w, c := range words {
		fmt.Printf("%s: %d\n", w, c)
	}
}
