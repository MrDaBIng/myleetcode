// Created by yangshow9226 at 2023/08/17 07:55
// leetgo: 1.3.3
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	if len(prices) < 2 {
		return 0
	}

	ans = 0
	for i := 1; i < len(prices); i++ {
		j := i - 1
		if prices[i] > prices[j] {
			ans = ans + prices[i] - prices[j]
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
