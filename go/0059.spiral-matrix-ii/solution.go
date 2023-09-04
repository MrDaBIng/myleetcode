// Created by yangshow9226 at 2023/09/02 15:06
// leetgo: 1.3.3
// https://leetcode.cn/problems/spiral-matrix-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateMatrix(n int) (ans [][]int) {
	ans = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	counter, offSet := 1, 1
	startI, startJ := 0, 0
	loop := n / 2
	for loop > 0 {
		i, j := startI, startJ
		for j = startJ; j < n-offSet; j++ {
			ans[i][j] = counter
			counter++
		}

		for i = startI; i < n-offSet; i++ {
			ans[i][j] = counter
			counter++
		}

		for ; j > startJ; j-- {
			ans[i][j] = counter
			counter++

		}

		for ; i > startI; i-- {
			ans[i][j] = counter
			counter++
		}

		offSet++
		startI++
		startJ++
		loop--
	}
	if n%2 == 1 {
		ans[n/2][n/2] = n * n
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateMatrix(n)

	fmt.Println("\noutput:", Serialize(ans))
}
