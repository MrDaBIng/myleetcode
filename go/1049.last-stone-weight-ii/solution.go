// Created by yangshow9226 at 2023/08/29 08:07
// leetgo: 1.3.3
// https://leetcode.cn/problems/last-stone-weight-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lastStoneWeightII(stones []int) (ans int) {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := range stones {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], stones[i]+dp[j-stones[i]])
		}
	}

	return sum - dp[target] - dp[target]
}

func max(a, b int) int {
	m := a
	if b > a {
		m = b
	}
	return m
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stones := Deserialize[[]int](ReadLine(stdin))
	ans := lastStoneWeightII(stones)

	fmt.Println("\noutput:", Serialize(ans))
}
