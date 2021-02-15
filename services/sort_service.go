package services

import (
	"github.com/gugazimmermann/xteam-bootcamp-go/sort"
)

// SortAscending - sorts a slice of ints in increasing order.
func SortAscending(elem []int) {
	sort.BubbleSortAscending(elem)
}

// SortDescending - sorts a slice of ints in decreasing order.
func SortDescending(elem []int) {
	sort.BubbleSortDescending(elem)
}

// Sort - sorts a slice of ints in increasing order.
func Sort(elem []int) {
	sort.Sort(elem)
}
