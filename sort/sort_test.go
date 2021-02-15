package sort

import (
	"testing"
	"time"
)

func getElem(n int) []int {
	res := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		res[j] = i
		j++
	}
	return res
}

// Basic Unit Test
func TestBubbleSortAscending(t *testing.T) {
	// Init:
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	// Execution:
	BubbleSortAscending(elem)
	// Validation:
	if elem[0] != 0 {
		t.Error("First element should be 0")
	}
	if elem[len(elem)-1] != 9 {
		t.Error("Last element should be 9")
	}
}

func TestBubbleSortAscending_AlreadySorted(t *testing.T) {
	elem := []int{0, 1, 2, 3, 4}
	BubbleSortAscending(elem)
	if elem[0] != 0 {
		t.Error("First element should be 0")
	}
	if elem[len(elem)-1] != 4 {
		t.Error("Last element should be 4")
	}
}

//Do not trust test or coverage if don't have validation
func TestBubbleSortDescending(t *testing.T) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	BubbleSortDescending(elem)
}

func TestBubbleSortDescending_AlreadySorted(t *testing.T) {
	elem := []int{0, 1, 2, 3, 4}
	BubbleSortDescending(elem)
}

func TestSort(t *testing.T) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	Sort(elem)
	if elem[0] != 0 {
		t.Error("First element should be 0")
	}
	if elem[len(elem)-1] != 9 {
		t.Error("Last element should be 9")
	}
}

func TestBubbleSortAscendingTimeOut(t *testing.T) {
	elem := getElem(10) // change to 100000 to timeout
	timeoutChan := make(chan bool, 1)
	go func() {
		BubbleSortAscending(elem)
		timeoutChan <- false
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()
	if <-timeoutChan {
		t.Error("took more than 500ms")
	}

}

func BenchmarkBubbleSortAscending(b *testing.B) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		BubbleSortAscending(elem)
	}
}

func BenchmarkSort(b *testing.B) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		Sort(elem)
	}
}

func BenchmarkBubbleSortAscending2(b *testing.B) {
	elem := getElem(100000)
	for i := 0; i < b.N; i++ {
		BubbleSortAscending(elem)
	}
}

func BenchmarkSort2(b *testing.B) {
	elem := getElem(100000)
	for i := 0; i < b.N; i++ {
		Sort(elem)
	}
}
