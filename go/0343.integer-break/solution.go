// Created by yangshow9226 at 2023/08/25 07:47
// leetgo: 1.3.3
// https://leetcode.cn/problems/integer-break/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func integerBreak(n int) (ans int) {
	if n < 2 {
		return -1
	}
	dp := make([]int, n+1)

	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
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
	n := Deserialize[int](ReadLine(stdin))
	ans := integerBreak(n)

	fmt.Println("\noutput:", Serialize(ans))
}
