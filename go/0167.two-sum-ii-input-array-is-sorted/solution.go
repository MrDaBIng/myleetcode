// Created by Show at 2023/07/10 14:24
// leetgo: 1.3.3
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// func twoSum(numbers []int, target int) (ans []int) {
// 	if len(numbers) <= 1 {
// 		return nil
// 	}
// 	for i := 0; i < len(numbers); i++ {
// 		for j := i + 1; j < len(numbers); j++ {
// 			if numbers[i]+numbers[j] == target {
// 				return []int{i + 1, j + 1}
// 			}
// 		}
// 	}
//
// 	return nil
// }

// @lc code=begin

func twoSum(numbers []int, target int) (ans []int) {
	if len(numbers) <= 1 {
		return nil
	}

	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		}
		if sum < target {
			i++
		}
	}

	return nil
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numbers := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(numbers, target)

	fmt.Println("\noutput:", Serialize(ans))
}
