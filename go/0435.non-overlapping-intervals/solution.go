// Created by yangshow9226 at 2023/08/15 07:58
// leetgo: 1.3.3
// https://leetcode.cn/problems/non-overlapping-intervals/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func eraseOverlapIntervals(intervals [][]int) (ans int) {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = 0
	for i := 1; i < len(intervals); i++ {
		j := i - 1
		if intervals[i][0] < intervals[j][1] {
			ans++
			intervals[i][1] = min(intervals[i][1], intervals[j][1])
		}
	}

	return ans
}

func min(a, b int) int {
	m := a
	if a > b {
		m = b
	}
	return m
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	ans := eraseOverlapIntervals(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
