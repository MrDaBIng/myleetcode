// Created by yangshow9226 at 2023/09/01 07:29
// leetgo: 1.3.3
// https://leetcode.cn/problems/house-robber/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rob(nums []int) (ans int) {
	// dp[i] 偷i号房间时的最高金额
	// 不偷i : dp[i-1]
	// 偷i: dp[i-2]+nums[i]
	// ==> dp[i]=max(dp[i-1], dp[i-2]+nums[i])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		fmt.Println(dp)
	}
	return dp[len(nums)-1]
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
