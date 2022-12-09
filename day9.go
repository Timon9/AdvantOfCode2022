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
	fmt.Println("Input:", input)

	var head, tail Coord
	visited := make(map[Coord]bool)

	// Iterate over each line in the input
	for _, line := range strings.Split(input, "\n") {
		var dir string
		var steps int

		// Parse the direction and number of steps
		fmt.Sscanf(line, "%1s%d", &dir, &steps)

		// Move the head the given number of steps
		for i := 0; i < steps; i++ {
			head.move(dir)

			// If the head is not adjacent to the tail, move the tail to the
			// same position as the head and mark the position as visited
			if head.x != tail.x && head.y != tail.y {
				if head.x > tail.x {
					for i := tail.x + 1; i < head.x; i++ {
						visited[Coord{x: i, y: head.y}] = true
					}
				} else if head.x < tail.x {
					for i := head.x + 1; i < tail.x; i++ {
						visited[Coord{x: i, y: head.y}] = true
					}
				} else if head.y > tail.y {
					for i := tail.y + 1; i < head.y; i++ {
						visited[Coord{x: head.x, y: i}] = true
					}
				} else if head.y < tail.y {
					for i := head.y + 1; i < tail.y; i++ {
						visited[Coord{x: head.x, y: i}] = true
					}
				}
			}
		}
	}

	// Count the number of unique positions visited by the tail
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}

	return count - 1
}

func solveDay9P1(input string) {
	count := countVisitedPositions(input)
	fmt.Println("countVisitedPositions:", count)

}
