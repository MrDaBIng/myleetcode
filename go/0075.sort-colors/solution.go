// Created by yangshow9226 at 2023/08/11 08:41
// leetgo: 1.3.3
// https://leetcode.cn/problems/sort-colors/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
var colors = []int{0, 1, 2}

func sortColors(nums []int) {
	bu := make(map[int][]int)
	for _, item := range nums {
		_, ok := bu[item]
		if !ok {
			bu[item] = []int{item}
			continue
		}
		bu[item] = append(bu[item], item)
	}

	i := 0
	for _, c := range colors {
		sub := bu[c]
		for _, item := range sub {
			nums[i] = item
			i++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	sortColors(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
