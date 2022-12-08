package main

import (
	"fmt"
	"testing"
)

func TestCanFindVisibleTrees(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	if r := findVisibleTrees(input); r != 21 {
		t.Fatalf("Part 1 FindVisibleTrees: Expected 21 but got %v", r)
	}
	fmt.Println(input)
}
