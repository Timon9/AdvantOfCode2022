package main

import (
	"fmt"
	"math"
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

// CalculateTrail calculates the trail of X and Y coordinates between two points.
// The start point is represented as (x1, y1) and the end point is represented as (x2, y2).
// The function returns a slice of coordinates representing the trail.
// NOTE! It need to START going the START POSITIONS of head before going to the END postion of head
func CalculateTrail(x1, y1, x2, y2 int) []Coord {
	fmt.Println("From ", x1, y1, "to", x2, y2)
	// Calculate the distance between the two points.
	dx := x2 - x1
	dy := y2 - y1

	// Calculate the number of steps needed to traverse the trail.
	// This is the maximum of the absolute values of the differences in x and y coordinates.
	steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))

	// Initialize a slice to hold the coordinates of the trail.
	trail := make([]Coord, steps+1)

	// Calculate the increment in x and y coordinates for each step.
	dxStep := dx / steps
	dyStep := dy / steps

	// Starting from the start point, calculate the coordinates of each point along the trail
	// by adding the corresponding increments to the previous coordinates.
	trail[0] = Coord{x: x1, y: y1}
	for i := 1; i <= steps; i++ {
		trail[i] = Coord{
			x: trail[i-1].x + dxStep,
			y: trail[i-1].y + dyStep,
		}
	}

	return trail
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
		fmt.Println("CalculateTrail", CalculateTrail(tail.x, tail.y, head.x, head.y))
		tail.x = head.x - 1
		tail.y = head.y
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
