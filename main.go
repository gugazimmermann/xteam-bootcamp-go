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

func (a *Animnal) eat() {
	fmt.Printf("The %v starts to %v and finds an %v\n", a.name, a.locomotion, a.food)
	fmt.Printf("The %v ate the %v! %v %v\n", a.name, a.food, a.noise, a.noise)
	a.hunger = false
}

func (a *Animnal) happy(s string) string {
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
	fmt.Println(bird.happy("EAT"))
	seeHunger(bird)
	bird.eat()
	seeHunger(bird)
	fmt.Println(bird.happy(""))

	fmt.Println()
	fmt.Println(snake.happy("TO DINNER"))
	seeHunger(snake)
	snake.eat()
	seeHunger(snake)
	fmt.Println(snake.happy(""))
}
