// Created by yangshow9226 at 2023/08/21 21:39
// leetgo: 1.3.3
// https://leetcode.cn/problems/unique-paths-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniquePathsWithObstacles(obstacleGrid [][]int) (ans int) {
	yLen := len(obstacleGrid)
	xLen := len(obstacleGrid[0])

	dp := make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		dp[i] = make([]int, xLen)
	}

	step := 1
	for i := 0; i < yLen; i++ {
		o := obstacleGrid[i][0]
		if o == 1 {
			step = 0
		}
		dp[i][0] = step

	}
	step = 1
	for i := 0; i < xLen; i++ {
		o := obstacleGrid[0][i]
		if o == 1 {
			step = 0
		}
		dp[0][i] = step
	}

	// dp[i][j]=dp[i-1][j]+dp[i][j-1]
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[yLen-1][xLen-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	obstacleGrid := Deserialize[[][]int](ReadLine(stdin))
	ans := uniquePathsWithObstacles(obstacleGrid)

	fmt.Println("\noutput:", Serialize(ans))
}
