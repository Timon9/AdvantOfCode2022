package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func findVisibleTrees(input string) int {
	return 0
}

func day8Part1(input string) {
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day8() {

	godotenv.Read(".env")

	fmt.Println("Day 8")

	input := GetInput(8)
	day8Part1(input)

}
