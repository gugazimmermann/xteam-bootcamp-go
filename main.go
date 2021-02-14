package main

import (
	"fmt"
)

func routine1(i int, c *int) {
	for x := 1; x <= 10; x++ {
		*c++
		fmt.Printf("Routine %d - Incremented value - now %d.\n", i, *c)
	}
}

func routine2(i int, c *int) {
	for x := 1; x <= 10; x++ {
		*c++
		fmt.Printf("Routine %d - Incremented value - now %d.\n", i, *c)
	}
}

// func routine2(i int, c *int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for x := 1; x <= 10; x++ {
// 		*c++
// 		fmt.Printf("Routine %d - Incremented value - now %d.\n", i, *c)
// 	}
// }

func main() {
	counter := 0
	routine1(1, &counter)
	routine1(2, &counter)
	fmt.Printf("Final Counter (should be 20): %d\n", counter)
	fmt.Println()

	counter = 0
	go routine2(1, &counter)
	go routine2(2, &counter)
	fmt.Printf("Final Counter (should be 20): %d\n", counter)

	// wg := &sync.WaitGroup{}
	// counter = 0
	// wg.Add(1)
	// go routine2(1, &counter, wg)
	// wg.Add(1)
	// go routine2(2, &counter, wg)
	// wg.Wait()
	// fmt.Printf("Final Counter (should be 2): %d\n", counter)
}
