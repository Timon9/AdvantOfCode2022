package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func day9Part1(input string) {
}

func day9Part2(input string) {
}

func printRightNumberLeftLetter(input string) {
	// Split input string into lines
	lines := strings.Split(input, "\n")

	headX := 0
	tailX := 0
	headY := 0

	for _, line := range lines {
		// Trim leading and trailing whitespace from line
		line = strings.TrimSpace(line)

		// Split line into words
		words := strings.Split(line, " ")

		// Get the right number and left letter
		rightNumber, _ := strconv.Atoi(string(words[len(words)-1]))
		leftLetter := words[0]

		switch string(leftLetter) {
		case "U":
			headY = headY - rightNumber
		case "D":
			headY = headY + rightNumber
		case "L":
			headX = headX - rightNumber
		case "R":
			headX = headX + rightNumber
		}

		tailToRight := (headX - tailX) - 1
		tailToLeft := (tailX - headX) - 1
		if tailToRight > 0 {
			fmt.Println("tailToRight", tailToRight)
			tailX = headX - 1
		}
		if tailToLeft > 0 {
			fmt.Println("tailToLeft", tailToLeft)
			tailX = headX + 1
		}

	}
}

func day9() {

	godotenv.Read(".env")

	fmt.Println("Day 9")

	input := GetInput(9)
	day9Part2(input)

}
