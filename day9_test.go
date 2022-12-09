package main

import (
	"fmt"
	"testing"
)

func TestDay9P1(t *testing.T) {

	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	result := countVisitedPositions(input, 0)
	fmt.Println("Result:", result) // Expected output: 36
}
