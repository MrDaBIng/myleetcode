// Created by yangshow9226 at 2023/08/20 22:42
// leetgo: 1.3.3
// https://leetcode.cn/problems/minimum-path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minPathSum(grid [][]int) (ans int) {
	yLen := len(grid)
	xLen := len(grid[0])
	dp := make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		dp[i] = make([]int, xLen)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < xLen; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	fmt.Println(dp)
	for i := 1; i < yLen; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	fmt.Println(dp)

	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[yLen-1][xLen-1]
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
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := minPathSum(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
