package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func day1() {

	godotenv.Read(".env")

	fmt.Println("Day 1")

	input := strings.Split(GetInput(), "\n") // Break the string into an array

	/**
	Solve the puzzle in O(n)
	**/
	c := 0 // Highest sub-count
	d := 0 // Second highest sub-count
	e := 0 // Third highest sub-count

	s := 0 // Counter per Elv
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			s = 0 // Reset the counter, an empty line indicates a new Elv
		} else {
			p, _ := strconv.Atoi(input[i]) // Convert string to int
			s = s + p
			if s > c { // Highest counter
				c = s
			} else if s > d { // Second highest counter
				d = s
			} else if s > e { // Third highest counter
				e = s
			}

		}
	}

	fmt.Println("Solution is:", (c + d + e))

}
