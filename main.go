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

	fmt.Println()
	fmt.Println()

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	sliceAppend(mySlice)
	fmt.Println(mySlice)
	sliceChange(mySlice)
	fmt.Println(mySlice)
}

/*
the contents of a slice argument can be modified by a function, but its header cannot.
The length stored in the slice variable is not modified by the call to the function,
since the function is passed a copy of the slice header, not the original.
Thus if we want to write a function that modifies the header, we must return it as a
result parameter, just as we have done here.
*/
func sliceAppend(a []int) {
	a = append(a, 4, 5, 6)
	fmt.Println(a)
}

func sliceChange(a []int) {
	a[2] = 9
}
