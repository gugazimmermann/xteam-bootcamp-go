package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	fmt.Println("Olá from main")
	wg.Wait()

	fmt.Println()

	quitSignal := make(chan bool)
	go sayHelloFromChannel(quitSignal)
	fmt.Println("Olá from main again")
	<-quitSignal

	fmt.Println()

	translatorSignal := make(chan string)

	lang := "german"
	go sayHelloInOtherLanguage(translatorSignal, lang, "Hello from a channel")
	fmt.Printf("Say in %v: ", lang)
	ts := <-translatorSignal
	fmt.Println("_", ts)

	lang = "spanish"
	go sayHelloInOtherLanguage(translatorSignal, lang, "Hello from a channel")
	fmt.Printf("Say in %v: ", lang)
	ts = <-translatorSignal
	fmt.Println("_", ts)

	fmt.Println()

	anyChan := make(chan bool)
	anyText := "Hello from anonymous function"
	go func() {
		fmt.Println(anyText)
		anyChan <- false
	}()
	<-anyChan

	fmt.Println()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from a new goroutine")
}

func sayHelloFromChannel(qs chan bool) {
	fmt.Println("Hola desde un canal")
	qs <- true
}

func sayHelloInOtherLanguage(ts chan string, lang, text string) {
	var translated string
	switch lang {
	case "german":
		trans := strings.NewReplacer("Hello", "Hallo", "from", "von", "a", "einem", "channel", "kanal")
		translated = trans.Replace(text)
	case "spanish":
		trans := strings.NewReplacer("Hello", "Hola", "from", "desde", "a", "un", "channel", "canal")
		translated = trans.Replace(text)
	default:
		translated = text
	}
	ts <- translated
}
