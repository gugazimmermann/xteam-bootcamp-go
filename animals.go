package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

var cow = animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}
var bird = animal{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}
var snake = animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.noise)
}

func info(t string, a animal) {
	switch t {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Information  not found")
	}
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	a := delete_empty(strings.Split(strings.TrimSpace(input), " "))
	switch a[0] {
	case "cow":
		info(a[1], cow)
	case "bird":
		info(a[1], bird)
	case "snake":
		info(a[1], snake)
	default:
		fmt.Println("Animal not found")
	}
}
