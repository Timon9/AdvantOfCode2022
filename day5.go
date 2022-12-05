package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func day5Part1(input []string) int {
	r := 0
	// Fill the matrix
	for i := 1; i < len(input) && i < 8; i++ {
		line := input[8-i]
		// crate1 := line[0:3]
		fmt.Println(line)
		// fmt.Println(crate1)

	}

	return r
}

func day5Part2(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {

	}

	return r
}

func day5() {

	godotenv.Read(".env")

	fmt.Println("Day 5")

	input := strings.Split(GetInput(5), "\n") // Break the string into an array
	fmt.Println("Part 1 solution is:", day5Part1(input))
	// fmt.Println("Part 2 solution is:", day5Part2(input))

}
