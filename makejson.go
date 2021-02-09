package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func formatString(x string) string {
	x = strings.Replace(x, "\n", "", -1)
	x = strings.Replace(x, "\r", "", -1)
	x = strings.ToLower(x)
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, type your name:")
	fmt.Print("-> ")
	n, _ := reader.ReadString('\n')
	fmt.Println("Please, type your address:")
	fmt.Print("-> ")
	a, _ := reader.ReadString('\n')
	m := map[string]string{"name": formatString(n), "address": formatString(a)}
	j, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(j))
	} else {
		fmt.Println("Error: ", err)
	}

}
