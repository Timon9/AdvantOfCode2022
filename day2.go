package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func day2() {

	godotenv.Read(".env")

	fmt.Println("Day 2")

	input := strings.Split(GetInput(2), "\n") // Break the string into an array

	/**
	Solve the puzzle in O(n)
	**/

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

	fmt.Println("Solution is:", c)

}
