package main

import "fmt"

func main() {
	s := "hi there"
	fmt.Printf("The string s is: %s\n", s)

	t := s
	fmt.Printf("The string t is: %s\n", t)

	s = s + "!"
	fmt.Printf("The string s modified with ! is: %s\n", s)
	fmt.Printf("The string t is: %s\n", t)
}
