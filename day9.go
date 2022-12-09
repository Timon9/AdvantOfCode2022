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

func (t *Coord) follow(h Coord, trail Coord) {
	if t.x < (h.x-1) || t.x > (h.x+1) || t.y < (h.y-1) || t.y > (h.y+1) {
		t.x = trail.x
		t.y = trail.y
	}
}

// countVisitedPositions counts the number of unique positions visited by the tail
// given a string of directions and steps
func countVisitedPositions(input string, extraKnots int) int {

	var head Coord
	tail := [12]Coord{}
	trail := [12]Coord{}
	//	var head, tail, trail Coord
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
			tail[0].follow(head, trail[0])
			trail[0] = head
			//
			visited[tail[9]] = true

		}
	}

	// Count the number of unique positions visited by the tail
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}

	return count
}

func solveDay9P1(input string) {
	result := countVisitedPositions(input, 0)
	fmt.Println("countVisitedPositions:", result)

}
