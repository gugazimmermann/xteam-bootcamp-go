package main

import "fmt"

type Cow struct {
	name, food, locomotion, noise string
	times                         int
}

type Bird struct {
	name, food, locomotion, noise string
	times                         int
}

type Snake struct {
	name, food, locomotion, noise string
	times                         int
}
type Animals interface {
	Name() string
	Eat()
	Move()
	Speak()
	Desc(s string)
	Times() int
	Inc()
}

func (a Cow) Name() string {
	return a.name
}
func (a *Cow) Times() int {
	return a.times
}
func (a *Cow) Eat() {
	fmt.Printf("found %v and ate\n", a.food)
}
func (a *Cow) Move() {
	fmt.Println("start to", a.locomotion)
}
func (a *Cow) Speak() {
	fmt.Println("is happy and speak: " + a.noise + "!")
}
func (a *Cow) Desc(str string) {
	fmt.Println(str)
}
func (a *Cow) Inc() {
	a.times++
}

func (a *Bird) Name() string {
	return a.name
}
func (a *Bird) Times() int {
	return a.times
}
func (a *Bird) Eat() {
	fmt.Printf("found a %v and ate\n", a.food)
}
func (a *Bird) Move() {
	fmt.Println("start to", a.locomotion)
}
func (a *Bird) Speak() {
	fmt.Println("is happy and speak: " + a.noise + "!")
}
func (a *Bird) Desc(str string) {
	fmt.Println(str)
}
func (a *Bird) Inc() {
	a.times++
}

func (a *Snake) Name() string {
	return a.name
}
func (a *Snake) Times() int {
	return a.times
}
func (a *Snake) Eat() {
	fmt.Printf("found a %v and ate\n", a.food)
}
func (a *Snake) Move() {
	fmt.Println("start to", a.locomotion)
}
func (a *Snake) Speak() {
	fmt.Println("is happy and speak: " + a.noise + "!")
}
func (a *Snake) Desc(str string) {
	fmt.Println(str)
}
func (a *Snake) Inc() {
	a.times++
}

func NewBird() *Bird {
	return &Bird{name: "Andorinha", food: "worms", locomotion: "flying", noise: "peep"}
}

func NewSnake(n string) Animals {
	s := new(Snake)
	s.name = n
	s.food = "mice"
	s.locomotion = "slither"
	s.noise = "hsss"
	return s
}

func main() {
	vaca := &Cow{name: "Vaca", food: "grass", locomotion: "walking", noise: "moo"}
	andorinha := NewBird()           //constructor
	cascavel := NewSnake("Cascavel") //constructor

	animals := []Animals{vaca, andorinha, cascavel}
	for _, a := range animals {
		a.Desc(fmt.Sprintf("%v is a animal!", a.Name()))
		fmt.Printf("%v eats %v times\n", a.Name(), a.Times())
		a.Move()
		a.Eat()
		a.Inc()
		fmt.Printf("%v eats %v times\n", a.Name(), a.Times())
		a.Speak()
		fmt.Println()
	}
}
