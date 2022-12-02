package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func part1(input []string) int {

	c := 0
	for i := 0; i < len(input); i++ {
		in := input[i]
		if len(in) < 2 {
			continue
		}
		a := input[i][0]
		b := input[i][2]

		if b == 'X' {
			c = c + 1
		} else if b == 'Y' {
			c = c + 2
		} else if b == 'Z' {
			c = c + 3
		}

		if a == 'A' {
			if b == 'X' {
				c = c + 3 // Draw
			} else if b == 'Y' {
				c = c + 6 // Win
			} else {
				c = c + 0 // Lost
			}
		}

		if a == 'B' {
			if b == 'X' {
				c = c + 0 // Lost
			} else if b == 'Y' {
				c = c + 3 // Draw
			} else {
				c = c + 6 // Won
			}
		}

		if a == 'C' {
			if b == 'X' {
				c = c + 6 // Win
			} else if b == 'Y' {
				c = c + 0 // Lost
			} else {
				c = c + 3 // Draw
			}
		}

	}
	return c
}

func part2(input []string) int {

	c := 0

	for i := 0; i < len(input); i++ {
		in := input[i]
		if len(in) < 2 {
			continue
		}
		a := input[i][0]
		b := input[i][2]
		switch b {
		case 'X': // Lose
			if a == 'A' { // Lose from rock: scisor
				fmt.Println("Losing from rock:scisor")
				c = c + 3
			} else if a == 'B' { // Lose from paper: rock
				fmt.Println("Losing from paper:scisor")
				c = c + 1
			} else if a == 'C' { // Lose from scisor: paper
				c = c + 2
			}
		case 'Y': // Draw
			c = c + 3
			if a == 'A' { // Draw from rock: rock
				c = c + 1
			} else if a == 'B' { // Draw from paper: paper
				c = c + 2
			} else if a == 'C' { // Draw from scisor: scisor
				c = c + 3
			}
		case 'Z': //Win
			c = c + 6
			if a == 'A' { // Win from rock: paper
				c = c + 2
			} else if a == 'B' { // Win from paper: scisor
				c = c + 3
			} else if a == 'C' { // Win from scisor: rock
				c = c + 1
			}
		}

	}
	return c
}

func day2() {

	godotenv.Read(".env")

	fmt.Println("Day 2")

	input := strings.Split(GetInput(2), "\n") // Break the string into an array

	fmt.Println("Part 1 solution is:", part1(input))
	fmt.Println("Part 2 solution is:", part2(input))

}
