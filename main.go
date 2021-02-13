package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("iptr:", iptr)
	fmt.Println("iptr dereference:", *iptr)
	*iptr = 0
}

type char struct {
	name string
}

func changeName1(c char) {
	c.name = "A2"
}

func changeName2(c *char) {
	c.name = "B2"
}

func changeName3(c *char) {
	c.name = "C2"
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("pointer:", &i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println()
	fmt.Println()

	char1 := new(char)
	char1.name = "A"
	fmt.Println("Character name: ", char1.name)
	changeName1(*char1)
	fmt.Println("Character name: ", char1.name)

	char2 := &char{"B"}
	fmt.Println("Character name: ", char2.name)
	changeName2(char2)
	fmt.Println("Character name: ", char2.name)

	char3 := char{"C"}
	fmt.Println("Character name: ", char3.name)
	changeName3(&char3)
	fmt.Println("Character name: ", char3.name)
}
