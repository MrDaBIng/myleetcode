// Created by yangshow9226 at 2023/08/28 08:41
// leetgo: 1.3.3
// https://leetcode.cn/problems/partition-equal-subset-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canPartition(nums []int) bool {
	sum := 0
	for idx := range nums {
		sum += nums[idx]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	// dp[j]=max(dp[j],nums[i]+dp[j-nums[i]])
	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], nums[i]+dp[j-nums[i]])
			if dp[j] == target {
				counter++
			}
		}
	}

	fmt.Println(dp)
	return (counter > 1)
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
	ans := canPartition(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
