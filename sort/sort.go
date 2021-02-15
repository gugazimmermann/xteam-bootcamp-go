package sort

import (
	"sort"
)

// BubbleSortAscending - sorts a slice of ints in increasing order.
func BubbleSortAscending(elem []int) {
	keepworking := true
	for keepworking {
		keepworking = false
		for i := 0; i < len(elem)-1; i++ {
			if elem[i] > elem[i+1] {
				keepworking = true
				elem[i], elem[i+1] = elem[i+1], elem[i]
			}
		}
	}
}

// BubbleSortDescending - sorts a slice of ints in decreasing order.
func BubbleSortDescending(elem []int) {
	keepworking := true
	for keepworking {
		keepworking = false
		for i := 0; i < len(elem)-1; i++ {
			if elem[i] < elem[i+1] {
				keepworking = true
				elem[i], elem[i+1] = elem[i+1], elem[i]
			}
		}
	}
}

// Sort - sorts a slice of ints in increasing order.
func Sort(elem []int) {
	sort.Ints(elem)
}
