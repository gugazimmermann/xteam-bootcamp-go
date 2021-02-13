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
}
