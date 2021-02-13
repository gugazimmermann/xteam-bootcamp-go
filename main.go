package main

import "fmt"

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
	name, food, locomotion, noise string
}

type Snake struct {
	name, food, locomotion, noise string
}
type Animals interface {
	Name() string
	Eat()
	Move()
	Speak()
}

func (c Cow) Name() string {
	return c.name
}
func (c Cow) Eat() {
	fmt.Println("eat: " + c.food)
}
func (c Cow) Move() {
	fmt.Println("move: " + c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println("speak: " + c.noise)
}

func (b Bird) Name() string {
	return b.name
}
func (b Bird) Eat() {
	fmt.Println("eat: " + b.food)
}
func (b Bird) Move() {
	fmt.Println("move: " + b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println("speak: " + b.noise)
}

func (s Snake) Name() string {
	return s.name
}
func (s Snake) Eat() {
	fmt.Println("eat: " + s.food)
}
func (s Snake) Move() {
	fmt.Println("move: " + s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println("speak: " + s.noise)
}

func main() {

	vaca := &Cow{name: "Vaca", food: "grass", locomotion: "walk", noise: "moo"}
	andorinha := &Bird{name: "Andorinha", food: "worms", locomotion: "fly", noise: "peep"}
	cascavel := &Snake{name: "Cascavel", food: "mice", locomotion: "slither", noise: "hsss"}
	animals := []Animals{vaca, andorinha, cascavel}
	for _, animal := range animals {
		fmt.Printf("%v\n", animal.Name())
		animal.Eat()
		animal.Move()
		animal.Speak()
		fmt.Println()
	}
}
