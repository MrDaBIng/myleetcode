// Created by yangshow9226 at 2023/09/01 08:39
// leetgo: 1.3.3
// https://leetcode.cn/problems/binary-search/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
			continue
		}
		if target < nums[mid] {
			right = mid - 1
			continue
		}
	}

	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
