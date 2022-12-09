package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func day9Part1(input string) {
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day9Part2(input string) {
	fmt.Println("Result part2", findBestTree(input))
}
func printRightNumberLeftLetter(input string) {
	// Split input string into lines
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Trim leading and trailing whitespace from line
		line = strings.TrimSpace(line)

		// Split line into words
		words := strings.Split(line, " ")

		// Get the right number and left letter
		rightNumber := words[len(words)-1]
		leftLetter := words[0]

		switch string(leftLetter) {
		case "U":
			fmt.Println("UP^", rightNumber)
		case "D":
			fmt.Println("DOWN_", rightNumber)
		case "L":
			fmt.Println("<LEFT", rightNumber)
		case "R":
			fmt.Println("RIGHT^", rightNumber)
		}
	}
}

func day9() {

	godotenv.Read(".env")

	fmt.Println("Day 9")

	input := GetInput(9)
	day9Part2(input)

}
