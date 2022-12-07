package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func findTotalSize(input string) int {

	lines := strings.Split(input, "\n")
	// cm := make(map[string]int) // Hashmap [dir] size

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) > 5 && line[0:5] == "$ cd " {
			fmt.Println("Chaging to...")
		}
	}
	return 0
}

func day7Part1(input string) {
	fmt.Println("Day7Part1")
}

func day7() {

	godotenv.Read(".env")

	fmt.Println("Day 7")

	input := GetInput(7)
	day7Part1(input)

}
