// Created by yangshow9226 at 2023/08/31 08:22
// leetgo: 1.3.3
// https://leetcode.cn/problems/perfect-squares/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSquares(n int) (ans int) {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j-i*i] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	if dp[n] == math.MaxInt {
		return -1
	}

	return dp[n]
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
	n := Deserialize[int](ReadLine(stdin))
	ans := numSquares(n)

	fmt.Println("\noutput:", Serialize(ans))
}
