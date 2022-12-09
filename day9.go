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

	// Iterate over each line in the input
	for _, line := range strings.Split(input, "\n") {
		var dir string
		var steps int

		// Parse the direction and number of steps
		fmt.Sscanf(line, "%1s%d", &dir, &steps)

		// Move the head the given number of steps
		for i := 0; i < steps; i++ {
			head.move(dir)
			//X
			if head.x > (tail.x + 1) {
				visited[Coord{x: (tail.x + 1), y: tail.y}] = true
				tail.x = tail.x + 1
			} else if head.x < (tail.x - 1) {
				visited[Coord{x: (tail.x - 1), y: tail.y}] = true
				tail.x = tail.x - 1
			}
			//Y
			if head.y > (tail.y + 1) {
				visited[Coord{x: tail.x, y: (tail.y + 1)}] = true
				tail.y = tail.y + 1
			} else if head.y < (tail.y - 1) {
				visited[Coord{x: tail.x, y: (tail.y - 1)}] = true
				tail.y = tail.y - 1
			}
			fmt.Println(tail)

			// // If the head is not adjacent to the tail, move the tail to the
			// // same position as the head and mark the position as visited
			// for i := tail.x; i < head.x-2; i++ {
			// 	visited[Coord{x: i, y: head.y}] = true
			// }
			// for i := head.x; i < tail.x-2; i++ {
			// 	visited[Coord{x: i, y: head.y}] = true
			// }
			// for i := tail.y; i < head.y-2; i++ {
			// 	visited[Coord{x: head.x, y: i}] = true
			// }
			// for i := head.y; i < tail.y-2; i++ {
			// 	visited[Coord{x: head.x, y: i}] = true
			// }

		}
		fmt.Println("------")

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
