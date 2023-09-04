// Created by yangshow9226 at 2023/09/02 10:53
// leetgo: 1.3.3
// https://leetcode.cn/problems/squares-of-a-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sortedSquares(nums []int) (ans []int) {
	ans = make([]int, len(nums))
	i, j, k := 0, len(nums)-1, len(nums)-1
	for i <= j {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			ans[k] = nums[i] * nums[i]
			i++
		} else {
			ans[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := sortedSquares(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
