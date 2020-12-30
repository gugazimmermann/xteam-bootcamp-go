package main

import (
	b "bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	var i string
	r := "Not Found!"
	fmt.Printf("Please, enter a string: ")
	scanner := b.NewScanner(os.Stdin)
	if scanner.Scan() {
		i = scanner.Text()
	}
	l := s.ReplaceAll(s.ToLower(i), " ", "")
	if s.HasPrefix(l, "i") && s.Contains(l, "a") && s.HasSuffix(l, "n") {
		r = "Found"
	}
	fmt.Printf(r)
}
