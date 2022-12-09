package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func day9Part1(input string) {
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day9Part2(input string) {
	fmt.Println("Result part2", findBestTree(input))
}

func day9() {

	godotenv.Read(".env")

	fmt.Println("Day 9")

	input := GetInput(9)
	day9Part2(input)

}
