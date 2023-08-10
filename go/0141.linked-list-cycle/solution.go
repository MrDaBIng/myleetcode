// Created by show at 2023/07/15 10:26
// leetgo: 1.3.3
// https://leetcode.cn/problems/linked-list-cycle/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	l1 := head
	l2 := head.Next
	for l1 != nil && l2 != nil && l2.Next != nil {
		if l1 == l2 {
			return true
		}
		l1 = l1.Next
		l2 = l2.Next.Next
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	_ = Deserialize[int](ReadLine(stdin))
	ans := hasCycle(head)

	fmt.Println("\noutput:", Serialize(ans))
}
