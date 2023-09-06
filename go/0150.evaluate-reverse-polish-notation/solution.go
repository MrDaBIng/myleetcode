// Created by show at 2023/09/06 23:38
// leetgo: 1.3.3
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
	"golang.org/x/text/number"
)

// @lc code=begin

func evalRPN(tokens []string) (ans int) {
	mmaps := map[string]func(_, _ int) int{
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

	stack := []string{}
	for _, t := range tokens {
		fn, ok := mmaps[t]
		if !ok {
			stack = append([]string{t}, stack...)
		} else {
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tokens := Deserialize[[]string](ReadLine(stdin))
	ans := evalRPN(tokens)

	fmt.Println("\noutput:", Serialize(ans))
}
