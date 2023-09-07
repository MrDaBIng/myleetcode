// Created by show at 2023/09/06 23:38
// leetgo: 1.3.3
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func evalRPN(tokens []string) (ans int) {
	mmaps := map[string]func(int, int) int{
		"+": func(num1, num2 int) int {
			return num1 + num2
		},
		"*": func(num1, num2 int) int {
			return num1 * num2
		},
		"-": func(num1, num2 int) int {
			return num1 - num2
		},
		"/": func(num1, num2 int) int {
			return num1 / num2
		},
	}

	stack := []int{}
	for _, t := range tokens {
		fn, ok := mmaps[t]
		if !ok {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		} else {
			length := len(stack)
			num2 := stack[length-1]
			stack = stack[:length-1]

			length = len(stack)
			num1 := stack[length-1]
			stack = stack[:length-1]

			stack = append(stack, fn(num1, num2))
		}
	}

	return stack[0]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tokens := Deserialize[[]string](ReadLine(stdin))
	ans := evalRPN(tokens)

	fmt.Println("\noutput:", Serialize(ans))
}
