package sorting_test

import (
	"fmt"
	"testing"

	"github.com/valteem/toona/sorting"
)

func TestInsertionSort(t *testing.T) {

	s := []int{11, 7, 34, 71, 3, 18, 31}

	f := func(a1, a2 int) bool {
		if a1 < a2 {
			return true
		} else {
			return false
		}
	}

	sorting.InsertionSort[int](s, f)

	fmt.Println(s)

	fmt.Println(sorting.InsertionSort[int]([]int{1}, f))

}