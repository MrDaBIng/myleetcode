// Created by yangshow9226 at 2023/09/05 07:35
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i <= j {
		exchange(i, j, s)
		i++
		j--
	}
}

func exchange(i, j int, s []byte) {
	s[i], s[j] = s[j], s[i]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[[]byte](ReadLine(stdin))
	reverseString(s)
	ans := s

	fmt.Println("\noutput:", Serialize(ans))
}
