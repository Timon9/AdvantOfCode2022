package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func findPriority(char rune) int {
	p := int(char - rune('A') + 1)
	if p >= 33 {
		p = p - 6
	}
	return p
}
func day3Part1(input []string) int {
	for i := 0; i < len(input); i++ {

	}
	return 0
}

func day3Part2(input []string) int {
	return 0
}

func day3() {

	godotenv.Read(".env")

	fmt.Println("Day 3")

	//input := strings.Split(GetInput(3), "\n") // Break the string into an array

	fmt.Println(findPriority('a'))
	fmt.Println(findPriority('z'))
	fmt.Println(findPriority('A'))
	fmt.Println(findPriority('Z'))
	//	fmt.Println("Part 1 solution is:", day3Part1(input))
	//	fmt.Println("Part 2 solution is:", day3Part2(input))

}
