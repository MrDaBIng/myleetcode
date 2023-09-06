// Created by yangshow9226 at 2023/09/06 08:04
// leetgo: 1.3.3
// https://leetcode.cn/problems/reverse-words-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseWords(s string) string {
	bts := []byte(s)
	fast, slow := 0, 0
	for ; fast < len(bts); fast++ {
		if bts[fast] != ' ' {
			if slow != 0 {
				bts[slow] = ' '
				slow++
			}
			for fast < len(bts) && bts[fast] != ' ' {
				bts[slow] = bts[fast]
				slow++
				fast++
			}
		}
	}

	bts = bts[0:slow]
	fmt.Println(string(bts))
	reverse(bts)
	fmt.Println(string(bts))

	last := 0
	for i := 0; i <= len(bts); i++ {
		if i == len(bts) || bts[i] == ' ' {
			reverse(bts[last:i])
			last = i + 1
		}
	}

	return string(bts)
}

func reverse(bts []byte) {
	i := 0
	j := len(bts) - 1
	for i < j {
		bts[i], bts[j] = bts[j], bts[i]
		i++
		j--
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseWords(s)

	fmt.Println("\noutput:", Serialize(ans))
}
