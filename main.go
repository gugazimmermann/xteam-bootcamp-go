package main

import (
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

func (a *Animnal) happy() string {
	msg := fmt.Sprintf("The %s is very happy :))", a.name)
	if a.hunger == true {
		msg = fmt.Sprintf("The %s is not happy >:(", a.name)
	}
	return msg
}

func main() {
	bird := &Animnal{"Bird", "earthworm", "fly", "peep", true}
	snake := &Animnal{"Snake", "mice", "slither", "hsss", true}

	fmt.Println()
	fmt.Println(bird.happy())
	seeHunger(bird)
	bird.eat()
	seeHunger(bird)
	fmt.Println(bird.happy())

	fmt.Println()
	fmt.Println(snake.happy())
	seeHunger(snake)
	snake.eat()
	seeHunger(snake)
	fmt.Println(snake.happy())
}
