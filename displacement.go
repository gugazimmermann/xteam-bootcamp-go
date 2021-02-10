package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func formatToFloat64(input string) float64 {
	inputFloat64, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}
	return inputFloat64
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		// ½ a t2 + vo t + so
		return (0.5*(a*(math.Pow(t, 2))) + (vo * t) + (so))
	}
	return fn
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, enter the value for acceleration:")
	fmt.Print("-> ")
	inputA, _ := reader.ReadString('\n')
	a := formatToFloat64(inputA)

	fmt.Println("Please, enter the value for initial velocity:")
	fmt.Print("-> ")
	inputVO, _ := reader.ReadString('\n')
	vo := formatToFloat64(inputVO)

	fmt.Println("Please, enter the value for initial displacement:")
	fmt.Print("-> ")
	inputSO, _ := reader.ReadString('\n')
	so := formatToFloat64(inputSO)

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println("Displacement after 3 seconds: ")
	fmt.Println(fn(3))

	fmt.Println("Displacement after 5 seconds: ")
	fmt.Println(fn(5))
}
