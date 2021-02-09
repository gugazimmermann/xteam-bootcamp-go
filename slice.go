package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func formatString(x string) string {
	x = strings.Replace(x, "\n", "", -1)
	x = strings.Replace(x, "\r", "", -1)
	x = strings.ToLower(x)
	return x
}

func main() {
	ints := make([]int, 3)
	s := ""
	x := 0
	for s != "x" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please, type a number (INT) or type X to exit:")
		fmt.Print("-> ")
		n, _ := reader.ReadString('\n')
		n = formatString(n)
		if strings.Compare("x", n) == 0 {
			s = n
			fmt.Println("Thanks!")
		} else {
			i, err := strconv.Atoi(n)
			if err == nil {
				if x < len(ints) {
					ints[0] = i
				} else {
					ints = append(ints, i)
				}
				x += 1
				sort.Ints(ints)
				fmt.Println("Ints:   ", ints)
			}
		}
	}
}
