package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
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

	ic := make(chan int)
	go periodicSend(ic)
	for i := range ic {
		fmt.Println(i)
	}
	// ic <- 6
	// panic: send on closed channel
	// with this we can see if channel are open or close
	_, ok := <-ic
	fmt.Println("Channel are open?", ok)

	fmt.Println()

	// bufferd channels block when they are either full or empty

	/* 	// Buffer 2, send 3 = deadlock
	   	buffc := make(chan int, 2)
	   	buffc <- 3 // FIFO
	   	buffc <- 2
	   	buffc <- 1
	   	fmt.Println(<-buffc)
	   	fmt.Println(<-buffc)
	*/
	/* 	// Buffer 5, have only 2, read 3 = deadlock
	   	buffc := make(chan int, 5)
	   	buffc <- 3 // FIFO
	   	buffc <- 2
	   	fmt.Println(<-buffc)
	   	fmt.Println(<-buffc)
	   	fmt.Println(<-buffc)
	*/
	buffc := make(chan int, 5)
	buffc <- 3 // FIFO
	buffc <- 2
	fmt.Println(<-buffc)
	fmt.Println(<-buffc)

	fmt.Println()

	// v2 return first, because wait less than v1

	testCh := make(chan int)
	select {
	case v1 := <-waitAndSend(1, 5):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1):
		fmt.Println(v2)
	case testCh <- 20:
		fmt.Println("testCh received a value")
		// default:
		// 	fmt.Println("All channels are slow")
	}
	// select can handle multiple channels, but testCh will
	// not be invoked because do not have a goroutine listening.

	// if default is not commented it will run first, because the
	// fastest channel still has to wait 1 second.
	// with default we can prevent select from blocking the code.

	// if two or more channel return at the same time
	// select will pick one at random.
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

func periodicSend(ic chan int) {
	i := 1
	for i <= 3 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	// without close() = fatal error: all goroutines are asleep - deadlock!
	close(ic)
}

func waitAndSend(v, t int) chan int {
	retCh := make(chan int)
	go func() {
		time.Sleep(time.Duration(t) * time.Second)
		retCh <- v
	}()
	return retCh
}
