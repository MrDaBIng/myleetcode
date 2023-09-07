// Created by show at 2023/09/07 22:52
// leetgo: 1.3.3
// https://leetcode.cn/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) (ans []int) {
	numMap := make(map[int]int)
	for idx := range nums {
		s := target - nums[idx]
		value, ok := numMap[s]
		if ok {
			return []int{idx, value}
		}
		numMap[nums[idx]] = idx
	}
	return []int{}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
