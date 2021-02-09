package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f float64
	fmt.Printf("Please, enter a floating point number: ")
	fmt.Scan(&f)
	fmt.Printf(strconv.Itoa(int(f)))
}
