package main

import (
	"strings"
	"testing"
)

func TestMoveOne(t *testing.T) {

	input := strings.Split(GetInput(5), "\n")

	matrix := make(map[int][]string)
	buildMatrix(input, matrix)

	for i := 10; i < len(input) && len(input[i]) > 4 && i <= 12; i++ {
		nr, from, to := decodeToFrom(input[i])
		matrix = moveV2(matrix, from, to, nr)

	}

	r := ""
	for i := 1; i <= len(matrix); i++ {
		s := len(matrix[i]) - 1
		if s > 0 {
			r = r + matrix[i][s]
		}
	}
	if r != "RBLMGVLD" {
		t.Fatalf("Expected `RBLMGVLD` but got %v", r)
	}
}
