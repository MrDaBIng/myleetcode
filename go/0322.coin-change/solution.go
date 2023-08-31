// Created by yangshow9226 at 2023/08/31 07:51
// leetgo: 1.3.3
// https://leetcode.cn/problems/coin-change/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func coinChange(coins []int, amount int) (ans int) {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] == math.MaxInt {
				continue
			}
			dp[j] = min(dp[j-coins[i]]+1, dp[j])
		}
	}

	if dp[amount] >= math.MaxInt {
		return -1
	}

	if dp[amount] < 0 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	coins := Deserialize[[]int](ReadLine(stdin))
	amount := Deserialize[int](ReadLine(stdin))
	ans := coinChange(coins, amount)

	fmt.Println("\noutput:", Serialize(ans))
}
