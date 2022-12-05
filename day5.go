package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func day5Part1(input []string) int {
	r := 0

	matrix := make(map[int][]string)

	// Fill the matrix
	for i := 1; i < len(input) && i < 8; i++ {
		line := input[8-i]
		crate1 := line[1:2]
		crate2 := line[5:6]
		crate3 := line[9:10]
		crate4 := line[13:14]
		crate5 := line[17:18]

		crate6 := line[21:22]
		crate7 := line[25:26]
		crate8 := line[29:30]
		crate9 := line[33:34]

		if crate1 != " " {
			matrix[1] = append(matrix[1], crate1)
		}
		if crate2 != " " {
			matrix[2] = append(matrix[2], crate2)
		}

		if crate3 != " " {
			matrix[3] = append(matrix[3], crate3)
		}
		if crate4 != " " {
			matrix[4] = append(matrix[4], crate4)
		}
		if crate5 != " " {
			matrix[5] = append(matrix[5], crate5)
		}
		if crate6 != " " {
			matrix[6] = append(matrix[6], crate6)
		}
		if crate7 != " " {
			matrix[7] = append(matrix[7], crate7)
		}
		if crate8 != " " {
			matrix[8] = append(matrix[8], crate8)
		}
		if crate9 != " " {
			matrix[9] = append(matrix[9], crate9)
		}

		fmt.Println(line)

	}
	fmt.Println(matrix)

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
