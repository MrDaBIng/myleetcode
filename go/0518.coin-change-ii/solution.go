// Created by yangshow9226 at 2023/08/30 08:07
// leetgo: 1.3.3
// https://leetcode.cn/problems/coin-change-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func change(amount int, coins []int) (ans int) {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}

	return dp[amount]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	amount := Deserialize[int](ReadLine(stdin))
	coins := Deserialize[[]int](ReadLine(stdin))
	ans := change(amount, coins)

	fmt.Println("\noutput:", Serialize(ans))
}
