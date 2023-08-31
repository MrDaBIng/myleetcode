// Created by yangshow9226 at 2023/08/29 08:37
// leetgo: 1.3.3
// https://leetcode.cn/problems/target-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findTargetSumWays(nums []int, target int) (ans int) {
	// 数组的总和为sum
	// 假设: 加法的和为x; 则减法的和为sum-x
	// target=sum-(sum-x) => x=(sum+target)/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 != 0 {
		return 0
	}
	x := (sum + target) / 2
	if x < 0 {
		return 0
	}

	dp := make([]int, x+1)
	dp[0] = 1

	// counter := 0
	for _, num := range nums {
		for j := x; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
			// if dp[j] == x {
			// 	counter++
			// }
		}
	}

	return dp[x]
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
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := findTargetSumWays(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
