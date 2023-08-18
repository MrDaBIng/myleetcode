// Created by yangshow9226 at 2023/08/18 07:50
// leetgo: 1.3.3
// https://leetcode.cn/problems/climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func climbStairs(n int) (ans int) {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := climbStairs(n)

	fmt.Println("\noutput:", Serialize(ans))
}
