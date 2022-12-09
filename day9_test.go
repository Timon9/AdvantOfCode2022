package main

import (
	"fmt"
	"testing"
)

func TestDay9P1(t *testing.T) {

	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	result := countVisitedPositions(input)
	fmt.Println("Result:", result) // Expected output: 13
}
