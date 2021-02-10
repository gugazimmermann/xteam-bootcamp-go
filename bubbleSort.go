package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type names struct {
	fname string
	lname string
}

func bubbleSort(intSlice []int) {
	for i := len(intSlice); i > 0; i-- {
		for j := 1; j < i; j++ {
			swap(intSlice, j)
		}
	}
	fmt.Println(intSlice)
}

func swap(intSlice []int, j int) {
	if intSlice[j-1] > intSlice[j] {
		intermediate := intSlice[j]
		intSlice[j] = intSlice[j-1]
		intSlice[j-1] = intermediate
	}
}

func formatToIntSlice(input string) []int {
	return upToTeen(unique(stringToInt(strings.Split(strings.Replace(strings.Replace(input, "\n", "", -1), "\r", "", -1), ","))))
}

func stringToInt(stringSlice []string) []int {
	list := []int{}
	for _, entry := range stringSlice {
		entryInt, _ := strconv.Atoi(entry)
		list = append(list, entryInt)
	}
	return list
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func upToTeen(intSlice []int) []int {
	if len(intSlice) > 10 {
		intSlice = intSlice[0:10]
	}
	return intSlice
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, type in a sequence of up to 10 integers...")
	fmt.Println("separated by commas (repeated numbers will be removed):")
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	intSlice := formatToIntSlice(input)
	bubbleSort(intSlice)
}
