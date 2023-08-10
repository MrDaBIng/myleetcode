// Created by yangshow9226 at 2023/08/09 08:00
// leetgo: 1.3.3
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findKthLargest(nums []int, k int) (ans int) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	// Sort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func partinting(nums []int, lo, hi int) int {
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

func Sort(nums []int, l, h int) {
	for l < h {
		p := partinting(nums, l, h)
		Sort(nums, l, p-1)
		Sort(nums, p+1, h)
	}
}

func exech(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
