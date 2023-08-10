// Created by yangshow9226 at 2023/08/10 08:16
// leetgo: 1.3.3
// https://leetcode.cn/problems/sort-characters-by-frequency/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func frequencySort(s string) string {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := frequencySort(s)

	fmt.Println("\noutput:", Serialize(ans))
}
