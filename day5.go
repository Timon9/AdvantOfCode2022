package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func move(matrix map[int][]string, from int, to int) map[int][]string {

	l := len(matrix[from]) - 1

	if l < 0 {
		return matrix
	}
	f := matrix[from][l]
	matrix[from] = matrix[from][0:l]
	matrix[to] = append(matrix[to], f)
	return matrix
}

func moveV2(matrix map[int][]string, from int, to int, r int) map[int][]string {

	for i := r; i > 0; i-- {
		l := len(matrix[from]) - i
		f := matrix[from][l]
		matrix[from] = append(matrix[from][0:l], matrix[from][l+1:len(matrix[from])]...)
		matrix[to] = append(matrix[to], f)
	}

	return matrix
}

func display(matrix map[int][]string) {
	for i := 1; i <= 9; i++ {
		fmt.Println(i, matrix[i])
	}

}

func buildMatrix(input []string, matrix map[int][]string) map[int][]string {

	// Fill the matrix
	for i := 1; i < len(input) && i <= 8; i++ {
		line := input[8-i]
		j := 1
		for c := 1; c <= 9; c++ {
			crate := line[j:(j + 1)]
			if crate != " " {
				fmt.Println("Building crate", i, crate)
				matrix[c] = append(matrix[c], crate)
			}
			j = j + 4
		}

	}
	return matrix
}

func day5Part1(input []string) {

	matrix := make(map[int][]string)
	buildMatrix(input, matrix)
	display(matrix)
	fmt.Println("=====---------====")
	for i := 10; i < len(input) && len(input[i]) > 4; i++ {
		if input[i][0:4] == "move" {

			l := strings.Split(input[i], "move ")
			nr, _ := strconv.Atoi(strings.Split(l[1], " ")[0])
			f := strings.Split(input[i], "from ")
			from, _ := strconv.Atoi(strings.Split(f[1], " ")[0])
			t := strings.Split(input[i], "to ")
			to, _ := strconv.Atoi(t[1])
			for ii := 1; ii <= nr; ii++ {
				matrix = move(matrix, from, to)
			}
		}

	}
	fmt.Println("=====---------====")

	display(matrix)
}

func day5Part2(input []string) {

	matrix := make(map[int][]string)
	buildMatrix(input, matrix)

	display(matrix)
	fmt.Println("=====---------====")
	for i := 10; i < len(input) && len(input[i]) > 4; i++ {
		if input[i][0:4] == "move" {

			l := strings.Split(input[i], "move ")
			nr, _ := strconv.Atoi(strings.Split(l[1], " ")[0])
			f := strings.Split(input[i], "from ")
			from, _ := strconv.Atoi(strings.Split(f[1], " ")[0])
			t := strings.Split(input[i], "to ")
			to, _ := strconv.Atoi(t[1])
			matrix = moveV2(matrix, from, to, nr)
		}

	}
	fmt.Println("=====---------====")

	display(matrix)
}

func day5() {

	godotenv.Read(".env")

	fmt.Println("Day 5")

	input := strings.Split(GetInput(5), "\n") // Break the string into an array
	day5Part2(input)

}
