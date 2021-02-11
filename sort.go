package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type names struct {
	fname string
	lname string
}

func formatToIntSlice(input string) []int {
	return stringToInt(strings.Split(strings.Replace(strings.Replace(input, "\n", "", -1), "\r", "", -1), ","))
}

func stringToInt(stringSlice []string) []int {
	list := []int{}
	for _, entry := range stringSlice {
		entryInt, _ := strconv.Atoi(entry)
		list = append(list, entryInt)
	}
	return list
}

func chunkSlice(slice []int) [][]int {
	var chunks [][]int
	start := 0
	n := len(slice) / 4
	r := len(slice) % 4
	for i := 0; i < 4; i++ {
		end := n + start
		if i < r {
			end++
		}
		chunks = append(chunks, slice[start:end])
		start = end
	}
	return chunks
}

func sortSlice(slice []int, wg *sync.WaitGroup) {
	fmt.Printf("%#v\n", slice)
	sort.Ints(slice)
	wg.Done()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, type in a sequence at least 4 integers separated by commas:")
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	intSlice := formatToIntSlice(input)
	for len(intSlice) < 4 {
		fmt.Println("Please, type in a sequence at least 4 integers separated by commas:")
		fmt.Print("-> ")
		input, _ = reader.ReadString('\n')
		intSlice = formatToIntSlice(input)
	}
	chunks := chunkSlice(intSlice)
	var wg sync.WaitGroup
	wg.Add(4)
	go sortSlice(chunks[0], &wg)
	go sortSlice(chunks[1], &wg)
	go sortSlice(chunks[2], &wg)
	go sortSlice(chunks[3], &wg)
	wg.Wait()
	var finalSlice []int
	for _, x := range chunks {
		finalSlice = append(finalSlice, x...)
	}
	fmt.Printf("%#v\n", finalSlice)
}
