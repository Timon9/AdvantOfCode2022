package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func findPriority(char rune) int {
	p := int(char) - 96
	if p < 0 {
		p = p + 58
	}
	return p
}

func findDuplicates(a string, b string) int {

	r := 0
	m := make(map[byte]int, 0)
	for i := 0; i < len(a); i++ {
		m[a[i]] = i // Fill the map
	}
	for ii := 0; ii < len(b); ii++ {
		if _, ok := m[b[ii]]; ok {

			a := findPriority(rune(b[ii]))
			// fmt.Println("Found", b[ii:ii+1], "Add", a)
			r = r + a
			delete(m, b[ii])
		}

	}
	return r
}

func day3Part1(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {
		sizeRucksack := len(input[i])
		com := sizeRucksack / 2
		r = r + findDuplicates(input[i][0:com], input[i][com:sizeRucksack])
	}
	return r
}

func day3Part2(input []string) int {
	return 0
}

func day3() {

	godotenv.Read(".env")

	fmt.Println("Day 3")

	input := strings.Split(GetInput(3), "\n") // Break the string into an array

	fmt.Println("Part 1 solution is:", day3Part1(input))
	//	fmt.Println("Part 2 solution is:", day3Part2(input))

}
