package main

import (
	"errors"
	"fmt"
)

type Animnal struct {
	name, food, locomotion, noise string
	hunger                        bool
}

func seeHunger(a *Animnal) {
	fmt.Printf("Is the %v hungry? %v\n", a.name, a.hunger)
}

func (a *Animnal) Eat() {
	fmt.Printf("The %v starts to %v and finds an %v\n", a.name, a.locomotion, a.food)
	fmt.Printf("The %v ate the %v! %v %v\n", a.name, a.food, a.noise, a.noise)
	a.hunger = false
}

func (a *Animnal) Happy(s string) string {
	if a.hunger == true {
		msg, err := say(s)
		if err == nil {
			fmt.Println(msg)
		}
		return fmt.Sprintf("The %s is not happy >:(", a.name)
	}
	return fmt.Sprintf("The %s is very happy :))", a.name)
}

func say(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	msg := fmt.Sprintf("I'M STUNNING WITH HUNGER! I WANNA %s", s)
	return msg, nil
}

func main() {
	bird := &Animnal{"Bird", "earthworm", "fly", "peep", true}
	snake := &Animnal{"Snake", "mice", "slither", "hsss", true}

	fmt.Println()
	fmt.Println(bird.Happy("EAT"))
	seeHunger(bird)
	bird.Eat()
	seeHunger(bird)
	fmt.Println(bird.Happy(""))

	fmt.Println()
	fmt.Println(snake.Happy("TO DINNER"))
	seeHunger(snake)
	snake.Eat()
	seeHunger(snake)
	fmt.Println(snake.Happy(""))

	fmt.Println()

	nuns := []int{1, 3, 5}
	variadic := Total(nuns...)
	fmt.Println(variadic)
}

//variadic
func Total(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
