package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func day7Part1(input string) {
	fmt.Println("Day7Part1")
}

func day7() {

	godotenv.Read(".env")

	fmt.Println("Day 7")

	input := GetInput(7)
	day7Part1(input)

}
