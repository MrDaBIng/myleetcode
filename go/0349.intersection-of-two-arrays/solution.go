// Created by yangshow9226 at 2023/09/07 08:43
// leetgo: 1.3.3
// https://leetcode.cn/problems/intersection-of-two-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func intersection(nums1 []int, nums2 []int) (ans []int) {
	mmap1 := make(map[int]struct{})
	mmap2 := make(map[int]struct{})
	for _, n := range nums1 {
		_, ok := mmap1[n]
		if !ok {
			mmap1[n] = struct{}{}
		}
	}
	for _, n := range nums2 {
		_, ok := mmap2[n]
		if !ok {
			mmap2[n] = struct{}{}
		}
	}

	for k := range mmap1 {
		_, ok := mmap2[k]
		if ok {
			ans = append(ans, k)
		}
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := intersection(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
