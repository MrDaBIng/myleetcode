package sort

import "fmt"

func partition(nums []int, lo, hi int) int {
	i := lo
	j := hi
	v := nums[lo]
	// 5, 4, 3, 5, 9, 8, 2
	// 0  1  2  3  4  5  6
	for {
		for nums[j] >= v && i < j {
			j--
		}

		for nums[i] <= v && i < j {
			i++
		}

		fmt.Println("i:", i, "j:", j)

		if i >= j {
			break
		}

		exech(nums, i, j)
		fmt.Println(nums)
	}

	exech(nums, lo, i)
	fmt.Println(nums)

	return i
}

func Partition(nums []int, lo, hi int) int {
	povid := nums[lo]
	i, j := lo, hi
	for {
		for nums[i] <= povid {
			i++
			if i >= hi {
				break
			}
		}
		for nums[j] >= povid {
			j--

			if j <= lo {
				break
			}
		}
		if i >= j {
			break
		}
		exech(nums, i, j)
	}

	exech(nums, j, lo)
	return j
}

func QuickSort2(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	povid := Partition(nums, lo, hi)
	QuickSort2(nums, lo, povid-1)
	QuickSort2(nums, povid+1, hi)
}

func exech(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func QuickSort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, lo, hi)
		QuickSort(nums, lo, p-1)
		QuickSort(nums, p+1, hi)
	}
}
