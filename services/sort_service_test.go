package services

import (
	"testing"
)

func TestSortAscending(t *testing.T) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	SortAscending(elem)
	if elem[0] != 0 {
		t.Error("First element should be 0")
	}
	if elem[len(elem)-1] != 9 {
		t.Error("Last element should be 9")
	}
}

func TestSortAscending_AlreadySorted(t *testing.T) {
	elem := []int{0, 1, 2, 3, 4}
	SortAscending(elem)
	if elem[0] != 0 {
		t.Error("First element should be 0")
	}
	if elem[len(elem)-1] != 4 {
		t.Error("Last element should be 4")
	}
}

func TestSortDescending(t *testing.T) {
	elem := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	SortDescending(elem)
	if elem[0] != 9 {
		t.Error("First element should be 9")
	}
	if elem[len(elem)-1] != 0 {
		t.Error("Last element should be 0")
	}
}

func TestSortDescending_AlreadySorted(t *testing.T) {
	elem := []int{9, 8, 7, 6}
	SortDescending(elem)
	if elem[0] != 9 {
		t.Error("First element should be 9")
	}
	if elem[len(elem)-1] != 6 {
		t.Error("Last element should be 6")
	}
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
