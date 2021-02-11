/*
A race condition is when two or more routines have access to the same resource concurrently,
such as a variable or data structure, and attempt to read and write to that resource without
any communication to the other routines.

This type of code can create the craziest and most random bugs you have ever seen.
It usually takes a tremendous amount of logging and luck to find these types of bugs.

You can run the program  with -race, this is show a report:

C:\Users\gugaz\Personal\xteam-bootcamp-go>go run -race racecondition.go
Routine 2 - Incremented value.
Routine 1 - Incremented value.
Routine 2 - Incremented value.
==================
WARNING: DATA RACE
Write at 0x0000009975e8 by goroutine 7:
  main.Routine()
      C:/Users/gugaz/Personal/xteam-bootcamp-go/racecondition.go:82 +0x106

Previous write at 0x0000009975e8 by goroutine 8:
  main.Routine()
      C:/Users/gugaz/Personal/xteam-bootcamp-go/racecondition.go:82 +0x106

Goroutine 7 (running) created at:
  main.main()
      C:/Users/gugaz/Personal/xteam-bootcamp-go/racecondition.go:64 +0x7c

Goroutine 8 (running) created at:
  main.main()
      C:/Users/gugaz/Personal/xteam-bootcamp-go/racecondition.go:64 +0x7c
==================
Routine 1 - Incremented value.
Final Counter (should be 4): 2
Found 1 data race(s)
exit status 66

*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup = sync.WaitGroup{}
var Counter int = 0
var Lock sync.Mutex = sync.Mutex{}

func main() {

	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter (should be 4): %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		// uncomment next line to fix the problem
		//Lock.Lock()

		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		fmt.Printf("Routine %d - Incremented value.\n", id)
		Counter = value

		// uncomment next line to fix the problem
		//Lock.Unlock()
	}

	Wait.Done()
}
