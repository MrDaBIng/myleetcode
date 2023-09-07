// Created by yangshow9226 at 2023/09/07 08:30
// leetgo: 1.3.3
// https://leetcode.cn/problems/valid-anagram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAnagram(s string, t string) bool {
	map1 := make(map[rune]int)

	for _, a := range s {
		map1[a]++
	}
	for _, a := range t {
		map1[a]--
	}

	for _, v := range map1 {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isAnagram(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
