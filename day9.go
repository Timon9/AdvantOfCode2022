package main

import (
	"fmt"
	"strings"
)

// Coord is a coordinate
type Coord struct {
	x, y int
}

func (c *Coord) move(dir string) {
	switch dir {
	case "R":
		c.x++
	case "L":
		c.x--
	case "U":
		c.y--
	case "D":
		c.y++
	}
}

// countVisitedPositions counts the number of unique positions visited by the tail
// given a string of directions and steps
func countVisitedPositions(input string) int {

	var head, tail Coord
	visited := make(map[Coord]bool)

	x := 0
	// Iterate over each line in the input
	for _, line := range strings.Split(input, "\n") {
		var dir string
		var steps int

		// Parse the direction and number of steps
		fmt.Sscanf(line, "%1s%d", &dir, &steps)

		// Move the head the given number of steps
		for i := 0; i < steps; i++ {
			head.move(dir)
		}
		/**
		Tail trails Head by 1, X/Y or Diagional
		*/
		// tail.x = head.x - 1
		// tail.y = head.y
		x++
		if x > 1 {
			return 0
		}
	}

	// Count the number of unique positions visited by the tail
	count := 1
	for _, v := range visited {
		if v {
			count++
		}
	}

	return count
}

func solveDay9P1(input string) {
	count := countVisitedPositions(input) //<45976
	fmt.Println("countVisitedPositions:", count)

}
