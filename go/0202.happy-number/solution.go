// Created by yangshow9226 at 2023/09/08 07:43
// leetgo: 1.3.3
// https://leetcode.cn/problems/happy-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isHappy(n int) bool {
	mapp := make(map[int]struct{})

	for n != 1 {
		n = getSum(n)
		_, ok := mapp[n]
		if ok {
			return false
		}
		mapp[n] = struct{}{}
	}

	return true
}

func getSum(num int) int {
	sum := 0

	for num != 0 {
		n := num % 10
		sum = sum + n*n
		num = num / 10
	}

	return sum
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := isHappy(n)

	fmt.Println("\noutput:", Serialize(ans))
}
