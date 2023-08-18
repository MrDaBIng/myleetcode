// Created by yangshow9226 at 2023/08/18 08:14
// leetgo: 1.3.3
// https://leetcode.cn/problems/fibonacci-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fib(n int) (ans int) {
	// F(0) = 0，F(1) = 1
	// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := fib(n)

	fmt.Println("\noutput:", Serialize(ans))
}
