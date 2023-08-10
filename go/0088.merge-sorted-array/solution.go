// Created by yangshow9226 at 2023/07/11 15:42
// leetgo: 1.3.3
// https://leetcode.cn/problems/merge-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(nums1 []int, m int, nums2 []int, n int) {
	mergeIndex := m + n - 1
	nums1Index := m - 1
	nums2Index := n - 1
	for nums2Index >= 0 {
		if nums1Index < 0 {
			nums1[mergeIndex] = nums2[nums2Index]
			mergeIndex--
			nums2Index--
			continue
		}
		if nums2Index < 0 {
			nums1[mergeIndex] = nums1[nums1Index]
			mergeIndex--
			nums1Index--
			continue
		}

		if nums1[nums1Index] >= nums2[nums2Index] {
			nums1[mergeIndex] = nums1[nums1Index]
			mergeIndex--
			nums1Index--
		} else {
			nums1[mergeIndex] = nums2[nums2Index]
			mergeIndex--
			nums2Index--
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	merge(nums1, m, nums2, n)
	ans := nums1

	fmt.Println("\noutput:", Serialize(ans))
}
