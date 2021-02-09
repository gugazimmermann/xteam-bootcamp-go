package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func formatString(x string) string {
	x = strings.Replace(x, "\n", "", -1)
	x = strings.Replace(x, "\r", "", -1)
	x = strings.ToLower(x)
	return x
}

func firstchars(s string, z int) string {
	runes := []rune(s)
	return string(runes[0:z])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, type the .txt file name:")
	fmt.Print("-> ")
	n, _ := reader.ReadString('\n')
	f, err := os.Open(formatString(n) + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sn := []*names{}
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		s := strings.Split(scn.Text(), " ")
		for i, x := range s {
			if len(x) > 2 {
				s[i] = firstchars(x, 20)
			}
		}
		ss := new(names)
		ss.fname = s[0]
		ss.lname = s[1]
		sn = append(sn, ss)
	}
	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	for _, x := range sn {
		fmt.Println(x.fname + " " + x.lname)
	}
}
