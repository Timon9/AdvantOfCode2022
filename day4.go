package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func day4Part1(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {

		pair := strings.Split(input[i], ",")
		if len(pair) > 1 {
			s1 := strings.Split(pair[0], "-")
			s2 := strings.Split(pair[1], "-")

			a1, _ := strconv.Atoi(s1[0])
			a2, _ := strconv.Atoi(s1[1])

			b1, _ := strconv.Atoi(s2[0])
			b2, _ := strconv.Atoi(s2[1])

			if a1 <= b1 && a2 >= b2 { // A overlaps B
				fmt.Println("1:", pair)
				r++
			} else if b1 <= a1 && b2 >= a2 { // B overlaps A
				fmt.Println("2:", pair)
				r++
			}
		}
	}

	return r
}

func day4Part2(input []string) int {
	r := 0
	for i := 0; i < len(input); i++ {

		pair := strings.Split(input[i], ",")
		if len(pair) > 1 {

			s1 := strings.Split(pair[0], "-")
			s2 := strings.Split(pair[1], "-")

			a1, _ := strconv.Atoi(s1[0])
			a2, _ := strconv.Atoi(s1[1])

			b1, _ := strconv.Atoi(s2[0])
			b2, _ := strconv.Atoi(s2[1])

			if (a1 < b1 && a2 < b1) || (b1 < a1 && b2 < a1) {
				fmt.Println("No overlap", pair)
			} else {
				r++
			}

		}
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
