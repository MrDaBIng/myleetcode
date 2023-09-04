// Created by yangshow9226 at 2023/09/02 11:14
// leetgo: 1.3.3
// https://leetcode.cn/problems/minimum-size-subarray-sum/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minSubArrayLen(target int, nums []int) (ans int) {
	i, j, sum := 0, 0, 0
	ans = math.MaxInt

	for ; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i-j+1)
			sum = sum - nums[j]
			j++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
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
	target := Deserialize[int](ReadLine(stdin))
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := minSubArrayLen(target, nums)

	fmt.Println("\noutput:", Serialize(ans))
}
