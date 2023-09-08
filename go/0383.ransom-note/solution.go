// Created by yangshow9226 at 2023/09/08 07:57
// leetgo: 1.3.3
// https://leetcode.cn/problems/ransom-note/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := make(map[rune]int)
	for _, b := range ransomNote {
		ransomNoteMap[b]++
	}
	fmt.Println(ransomNoteMap)
	magazineMap := make(map[rune]int)
	for _, b := range magazine {
		magazineMap[b]++
	}
	fmt.Println(magazineMap)

	if len(ransomNoteMap) > len(magazineMap) {
		return false
	}

	for rk, rv := range ransomNoteMap {
		mv, ok := magazineMap[rk]
		if !ok {
			return false
		}
		if rv > mv {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ransomNote := Deserialize[string](ReadLine(stdin))
	magazine := Deserialize[string](ReadLine(stdin))
	ans := canConstruct(ransomNote, magazine)

	fmt.Println("\noutput:", Serialize(ans))
}
