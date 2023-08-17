// Created by yangshow9226 at 2023/08/17 08:02
// leetgo: 1.3.3
// https://leetcode.cn/problems/queue-reconstruction-by-height/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reconstructQueue(people [][]int) (ans [][]int) {
	if len(people) < 2 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[j][0] < people[i][0]
	})
	fmt.Println("sort:", people)

	ans = [][]int{}
	for i := 0; i < len(people); i++ {
		idx := people[i][1]
		// insert to people by index
		if idx+1 > len(ans) {
			ans = append(ans, []int{})
			ans[idx] = people[i]
			continue
		}
		ans = append(ans[:idx+1], ans[idx:]...)
		// ans = append(ans, people[i])
		// for i := len(ans) - 1; i > idx; i-- {
		// 	ans[i] = ans[i-1]
		// }
		ans[idx] = people[i]
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	people := Deserialize[[][]int](ReadLine(stdin))
	ans := reconstructQueue(people)

	fmt.Println("\noutput:", Serialize(ans))
}
