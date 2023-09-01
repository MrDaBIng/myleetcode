// Created by yangshow9226 at 2023/09/01 07:47
// leetgo: 1.3.3
// https://leetcode.cn/problems/house-robber-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rob(nums []int) (ans int) {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	doRob := func(nums []int) int {
		length := len(nums)
		if length == 0 {
			return 0
		}
		if length == 1 {
			return nums[0]
		}

		dp := make([]int, length)
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])

		for i := 2; i < length; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
			fmt.Println(dp)
		}
		return dp[length-1]
	}

	// fmt.Println(nums[1:])
	// fmt.Println(nums[:length-1])
	res1 := doRob(nums[1:])
	res2 := doRob(nums[:length-1])
	return max(res1, res2)
}

func doRob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(dp[0], dp[1])

	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		fmt.Println(dp)
	}
	return dp[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := rob(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
