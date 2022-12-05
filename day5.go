package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func day5part1(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {
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

	// input := strings.Split(GetInput(5), "\n") // Break the string into an array
	input := []string{"123"}
	fmt.Println("Part 1 solution is:", day4Part1(input))
	// fmt.Println("Part 2 solution is:", day4Part2(input))

}
