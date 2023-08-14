// Created by yangshow9226 at 2023/08/15 07:42
// leetgo: 1.3.3
// https://leetcode.cn/problems/assign-cookies/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)

	g_idx, s_idx, ans := 0, 0, 0
	for g_idx < len(g) && s_idx < len(s) {
		if g[g_idx] <= s[s_idx] {
			ans++
			g_idx++
		}
		s_idx++
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	g := Deserialize[[]int](ReadLine(stdin))
	s := Deserialize[[]int](ReadLine(stdin))
	ans := findContentChildren(g, s)

	fmt.Println("\noutput:", Serialize(ans))
}
