// Created by yangshow9226 at 2023/08/17 22:55
// leetgo: 1.3.3
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	if len(prices) < 1 {
		return 0
	}
	ans = 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > ans {
				ans = prices[j] - prices[i]
			}
		}
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
