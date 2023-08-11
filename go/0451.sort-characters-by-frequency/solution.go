// Created by yangshow9226 at 2023/08/10 08:16
// leetgo: 1.3.3
// https://leetcode.cn/problems/sort-characters-by-frequency/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func frequencySort(s string) string {
	charaCounter := make(map[string]int)
	for _, c := range s {
		_, ok := charaCounter[string(c)]
		if !ok {
			charaCounter[string(c)] = 1
			continue
		}
		charaCounter[string(c)]++
	}
	queue := &Queue{}
	for s, count := range charaCounter {
		heap.Push(queue, &Item{
			chara: s,
			count: count,
		})
	}

	str := ""
	for queue.Len() > 0 {
		pop := heap.Pop(queue)
		item := pop.(*Item)
		for i := 0; i < item.count; i++ {
			str = str + item.chara
		}
	}

	return str
}

type Item struct {
	chara string
	count int
}

// Len is the number of elements in the collection.
func (q *Queue) Len() int {
	return len(*q)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (q *Queue) Less(i int, j int) bool {
	return (*q)[j].count < (*q)[i].count
}

// Swap swaps the elements with indexes i and j.
func (q *Queue) Swap(i int, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *Queue) Push(x any) {
	item, ok := x.(*Item)
	if ok {
		(*q) = append(*q, item)
	}
}

func (q *Queue) Pop() any {
	n := len(*q)
	item := (*q)[n-1]
	(*q) = (*q)[:n-1]
	return item
}

type Queue []*Item

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := frequencySort(s)

	fmt.Println("\noutput:", Serialize(ans))
}
