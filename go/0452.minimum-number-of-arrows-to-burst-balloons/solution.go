// Created by yangshow9226 at 2023/08/16 07:59
// leetgo: 1.3.3
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMinArrowShots(points [][]int) (ans int) {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans = 0
	for i := 1; i < len(points); i++ {
		j := i - 1
		if points[i][0] <= points[j][1] {
			ans++
			points[i][1] = min(points[i][1], points[j][1])
		}
	}

	return len(points) - ans
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
	points := Deserialize[[][]int](ReadLine(stdin))
	ans := findMinArrowShots(points)

	fmt.Println("\noutput:", Serialize(ans))
}
