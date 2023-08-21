// Created by yangshow9226 at 2023/08/19 10:38
// leetgo: 1.3.3
// https://leetcode.cn/problems/unique-paths/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniquePaths(m int, n int) (ans int) {
	// dp[m,n]=dp[m-1,n]+dp[m,n-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		slice := make([]int, n)
		slice[0] = 1
		dp[i] = slice
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// for i := 0; i < m; i++ {
	// 	dp[i][0] = 1
	// }
	fmt.Println(dp)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := uniquePaths(m, n)

	fmt.Println("\noutput:", Serialize(ans))
}
