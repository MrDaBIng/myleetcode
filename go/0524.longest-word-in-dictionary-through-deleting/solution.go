// Created by show at 2023/07/15 10:32
// leetgo: 1.3.3
// https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findLongestWord(s string, dictionary []string) string {
	str := ""
	for _, d := range dictionary {
		if isSebString(s, d) {
			if len(d) > len(str) {
				str = d
				continue
			}
			if len(d) == len(str) && Compare(d, str) {
				str = d
				continue
			}
		}
	}
	return str
}

// <  : -1
// == :  0
// >  :  1
func Compare(str1, str2 string) bool {
	for idx := range str1 {
		if str1[idx] > str2[idx] {
			return false
		}
	}

	return true
}

func isSebString(str, sub string) bool {
	if len(sub) > len(str) {
		return false
	}

	i := 0
	j := 0
	found := ""
	for i < len(sub) && j < len(str) {
		if str[j] == sub[i] {
			found = found + string(str[j])
			j++
			i++
			continue
		}
		j++
	}

	return found == sub
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	dictionary := Deserialize[[]string](ReadLine(stdin))
	ans := findLongestWord(s, dictionary)

	fmt.Println("\noutput:", Serialize(ans))
}
