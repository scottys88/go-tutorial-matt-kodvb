package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	fmt.Println("Enter text:")

	// buffer scans one word at a time
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		word := scan.Text()
		words[word]++
	}

	fmt.Println(len(words), "words found")
}
