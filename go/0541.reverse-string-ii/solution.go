// Created by yangshow9226 at 2023/09/05 07:46
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-string-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseStr(s string, k int) string {
	bs := []byte(s)
	length := len(bs)

	for i := 0; i < length; i = i + 2*k {
		if i+k < length {
			reverseBytes(i, i+k-1, bs)
			continue
		}
		reverseBytes(i, length-1, bs)
	}
	return string(bs)
}

func reverseBytes(start, end int, s []byte) {
	i, j := start, end
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := reverseStr(s, k)

	fmt.Println("\noutput:", Serialize(ans))
}
