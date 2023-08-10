package main

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	Sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
