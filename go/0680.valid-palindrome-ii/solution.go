// Created by yangshow9226 at 2023/07/11 08:44
// leetgo: 1.3.3
// https://leetcode.cn/problems/valid-palindrome-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	h := len(s) - 1
	l := 0
	for l < h {
		if s[h] != s[l] {
			return isValid(s[l+1:h+1]) || isValid(s[l:h])
		}
		l++
		h--
	}
	return true
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return true
	}
	low := 0
	high := len(s) - 1

	for low < high {
		if s[low] != s[high] {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := validPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
