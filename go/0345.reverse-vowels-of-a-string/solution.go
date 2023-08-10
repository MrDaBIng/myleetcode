// Created by yangshow9226 at 2023/07/10 15:44
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-vowels-of-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseVowels(s string) string {
	vowelsMap := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
		"A": {},
		"E": {},
		"I": {},
		"O": {},
		"U": {},
	}

	if len(s) <= 1 {
		return s
	}

	strBytes := []byte(s)
	l := 0
	r := len(strBytes) - 1
	for l < r {
		_, okL := vowelsMap[string(strBytes[l])]
		_, okR := vowelsMap[string(strBytes[r])]
		if !okL {
			l++
		}
		if !okR {
			r--
		}
		if okL && okR {
			tmp := strBytes[l]
			strBytes[l] = strBytes[r]
			strBytes[r] = tmp
		}
	}
	return string(strBytes)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseVowels(s)

	fmt.Println("\noutput:", Serialize(ans))
}
