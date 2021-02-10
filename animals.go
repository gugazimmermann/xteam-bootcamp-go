package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
	name, food, locomotion, noise string
}

type Snake struct {
	name, food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println(s.noise)
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

func nameExists(a []interface{}, n string) bool {
	for _, x := range a {
		cow, ok := x.(Cow)
		if ok {
			if cow.name == n {
				return true
			}
		}
		bird, ok := x.(Bird)
		if ok {
			if bird.name == n {
				return true
			}
		}
		snake, ok := x.(Snake)
		if ok {
			if snake.name == n {
				return true
			}
		}
	}
	return false
}

func createNew(a []interface{}, n, t string) interface{} {
	if nameExists(a, n) {
		fmt.Println("The name already exists, please choose a different one")
		return nil
	}
	switch t {
	case "cow":
		return Cow{name: n, food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		return Bird{name: n, food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		return Snake{name: n, food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		fmt.Println("Sorry, wrong animal type, please choose: cow, bird or snake")
		return nil
	}
}

func queryCow(c Cow, i string) {
	switch i {
	case "eat":
		c.Eat()
	case "move":
		c.Move()
	case "speak":
		c.Speak()
	default:
		fmt.Println("Sorry, wrong information type, please choose: eat, move or speak")
	}
}

func queryBird(b Bird, i string) {
	switch i {
	case "eat":
		b.Eat()
	case "move":
		b.Move()
	case "speak":
		b.Speak()
	default:
		fmt.Println("Sorry, wrong information type, please choose: eat, move or speak")
	}
}

func querySnake(s Snake, i string) {
	switch i {
	case "eat":
		s.Eat()
	case "move":
		s.Move()
	case "speak":
		s.Speak()
	default:
		fmt.Println("Sorry, wrong information type, please choose: eat, move or speak")
	}
}

func queryAnimal(a []interface{}, n, i string) {
	if !nameExists(a, n) {
		fmt.Println("Sorry, no animal with name: ", n)
	}
	for _, x := range a {
		cow, ok := x.(Cow)
		if ok {
			if cow.name == n {
				queryCow(cow, i)
			}
		}
		bird, ok := x.(Bird)
		if ok {
			if bird.name == n {
				queryBird(bird, i)
			}
		}
		snake, ok := x.(Snake)
		if ok {
			if snake.name == n {
				querySnake(snake, i)
			}
		}
	}
}

func main() {
	var animals []interface{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("")
		fmt.Println("If you want a new animal (type: cow|bird|snake): newanimal name type")
		fmt.Println("If you want a to query animal (information: eat|move|speak): query name information")
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		inputArr := delete_empty(strings.Split(strings.TrimSpace(input), " "))
		switch inputArr[0] {
		case "newanimal":
			newAnimal := createNew(animals, inputArr[1], inputArr[2])
			if newAnimal != nil {
				animals = append(animals, newAnimal)
				fmt.Println("Created it!")
			}
		case "query":
			queryAnimal(animals, inputArr[1], inputArr[2])
		default:
			fmt.Println("Sorry, wrong input type, please choose: newanimal or query")
		}
	}
}
