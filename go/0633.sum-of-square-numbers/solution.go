// Created by yangshow9226 at 2023/07/10 15:26
// leetgo: 1.3.3
// https://leetcode.cn/problems/sum-of-square-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	i := 0
	j := c
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		}
		if sum > c {
			j--
		} else {
			i++
		}
	}

	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	c := Deserialize[int](ReadLine(stdin))
	ans := judgeSquareSum(c)

	fmt.Println("\noutput:", Serialize(ans))
}
