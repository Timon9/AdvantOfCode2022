package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func findVisibleTrees(input string) int {
	r := 0
	lines := strings.Split(input, "\n")
	c := len(lines)
	a := 0
	for i := 0; i < len(lines[0]); i++ {
		a++
	}

	out := c + (a - 1) + (a - 1) + (c - 2)

	for i := 1; i < len(lines)-1; i++ {
		for ii := 1; ii < len(lines[i])-1; ii++ {
			fmt.Printf("%v,", string(lines[i][ii]))
		}

	}
	r = out
	return r
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
