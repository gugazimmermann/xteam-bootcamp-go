package main

import (
	"errors"
	"fmt"
	"log"
)

func seeError(b bool) (string, error) {
	if !b {
		return "", errors.New("b is false")
	}
	return "b is true", nil
}

type cat struct {
	name string
}
type dog struct {
	name string
}
type snake struct {
	name string
}

func main() {
	verdadeiro := true
	if verdadeiro {
		fmt.Println("verdadeiro in pt = true in en")
	} else {
		fmt.Println("falso in pt = false in en")
	}

	verdadeiro = false
	if !verdadeiro {
		fmt.Println("falso in pt = false in en")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	haverror, err := seeError(verdadeiro)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(haverror)

	haverror, err = seeError(true)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(haverror)

	i := 1
	for i <= 3 {
		fmt.Print(i)
		i = i + 1
	}
	fmt.Println()

	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := len(j) - 1; i >= 0; i-- {
		fmt.Print(j[i])
	}
	fmt.Println()

	for {
		fmt.Println("inf loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
	}
	fmt.Println()

	i = 2
	fmt.Printf("%d in portuguese is ", i)
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	color := "white"
	switch color {
	case "black", "white":
		fmt.Println("colorless")
	default:
		fmt.Println("colorful")
	}

	switch {
	case i < 10:
		fmt.Println("One digit")
	default:
		fmt.Println("More than one digit")
	}

	whichAnimal := func(i interface{}) {
		switch t := i.(type) {
		case cat:
			fmt.Println("Cat")
		case dog:
			fmt.Println("Dog")
		default:
			fmt.Printf("%T is unknow to me\n", t)
		}
	}
	cat1 := cat{"cat"}
	dog1 := dog{"cat"}
	snake1 := snake{"snake"}
	whichAnimal(cat1)
	whichAnimal(dog1)
	whichAnimal(snake1)
	fmt.Println()
	sum := 0
	for _, num := range j {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range j {
		if num == 3 {
			fmt.Println("num 3 index: ", i)
		}
	}

	m := map[int]string{1: "GO", 2: "GOLANG", 3: "GOPHER"}
	for i, v := range m {
		fmt.Printf("%d -> %s\n", i, v)
	}

	for k := range m {
		fmt.Println("key:", k)
	}

	for i, c := range "gopher" {
		fmt.Println(i, c)
	}
}
