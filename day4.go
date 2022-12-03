package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func day4Part1(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {
	}
	return r
}

func day4Part2(input []string) int {
	r := 0

	for i := 0; i < len(input); i++ {
	}

	return r
}

func day4() {

	godotenv.Read(".env")

	fmt.Println("Day 4")

	input := strings.Split(GetInput(4), "\n") // Break the string into an array

	fmt.Println("Part 1 solution is:", day4Part1(input))
	fmt.Println("Part 2 solution is:", day4Part2(input))

}
