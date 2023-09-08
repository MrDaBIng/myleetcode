package main

import (
	"fmt"
	"testing"
)

func Test_Name(t *testing.T) {
	ans := canConstruct("aab", "baa")
	fmt.Println(ans)
}
