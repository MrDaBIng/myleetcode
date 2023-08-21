// Created by yangshow9226 at 2023/08/18 08:26
// leetgo: 1.3.3
// https://leetcode.cn/problems/min-cost-climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minCostClimbingStairs(cost []int) (ans int) {
	dp := make([]int, len(cost))
	// cost[i] = min(dp[i-1],dp[i-2])+cost[i]

	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	m := a
	if a > b {
		m = b
	}
	return m
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := minCostClimbingStairs(cost)

	fmt.Println("\noutput:", Serialize(ans))
}
