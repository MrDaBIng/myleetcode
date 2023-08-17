package main

import (
	"fmt"
	"testing"
)

func TestQuene(t *testing.T) {
	peoples := [][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}}
	ans := reconstructQueue(peoples)
	fmt.Println(ans)
}

//  ● Expected:   [[3,0],[6,0],[7,0],[5,2],[3,4],[5,3],[6,2],[2,7],[9,0],[1,9]]
