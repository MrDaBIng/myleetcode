// Created by yangshow9226 at 2023/08/10 07:42
// leetgo: 1.3.3
// https://leetcode.cn/problems/top-k-frequent-elements/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func topKFrequent(nums []int, k int) (ans []int) {
	countor := make(map[int]int)
	for _, n := range nums {
		_, ok := countor[n]
		if !ok {
			countor[n] = 1
			continue
		}
		countor[n]++
	}

	pq := &PriorityQueue{}
	for num, count := range countor {
		heap.Push(pq, &Item{
			count: count,
			num:   num,
		})
	}

	ans = make([]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(pq)
		pop := item.(*Item)
		ans[i] = pop.num
	}

	return ans
}

type Item struct {
	num   int
	count int
}

type PriorityQueue []*Item

// Len is the number of elements in the collection.
func (q *PriorityQueue) Len() int {
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
func (q *PriorityQueue) Less(i int, j int) bool {
	return (*q)[j].count < (*q)[i].count
}

// Swap swaps the elements with indexes i and j.
func (q *PriorityQueue) Swap(i int, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PriorityQueue) Push(x any) {
	item, ok := x.(*Item)
	if ok {
		*q = append(*q, item)
	}
}

func (q *PriorityQueue) Pop() any {
	l := len(*q)
	item := (*q)[l-1]
	(*q) = (*q)[:l-1]
	return item
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := topKFrequent(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
