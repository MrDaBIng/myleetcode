// Created by show at 2023/09/07 23:27
// leetgo: 1.3.3
// https://leetcode.cn/problems/4sum-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (ans int) {
	mapp := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			mapp[n1+n2]++
		}
	}

	count := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			count += mapp[0-(n3+n4)]
		}
	}

	return count
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	nums3 := Deserialize[[]int](ReadLine(stdin))
	nums4 := Deserialize[[]int](ReadLine(stdin))
	ans := fourSumCount(nums1, nums2, nums3, nums4)

	fmt.Println("\noutput:", Serialize(ans))
}
