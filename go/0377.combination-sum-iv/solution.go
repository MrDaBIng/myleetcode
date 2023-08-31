// Created by yangshow9226 at 2023/08/30 08:16
// leetgo: 1.3.3
// https://leetcode.cn/problems/combination-sum-iv/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func combinationSum4(nums []int, target int) (ans int) {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := combinationSum4(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
