package main

import (
	"fmt"
	"testing"
)

func TestBestTree(t *testing.T) {
	in := `
30373
25512
65332
33549
35390
`

	fmt.Println(findBestTree(in))
}
