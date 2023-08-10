package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	num := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}

	QuickSort2(num, 0, len(num)-1)
	fmt.Println("res:", num)
}
