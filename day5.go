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

func day5Part1(input []string) {

	matrix := make(map[int][]string)

	// Fill the matrix
	for i := 1; i < len(input) && i <= 8; i++ {
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

	}
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
			// fmt.Println(nr, "x ", from, "=>", to)
			for ii := 1; ii <= nr; ii++ {
				matrix = move(matrix, from, to)
			}
		}
		// display(matrix)
		// fmt.Println("====")

	}
	fmt.Println("=====---------====")

	display(matrix)
}

func day5Part2(input []string) {

	matrix := make(map[int][]string)

	// Fill the matrix
	for i := 1; i < len(input) && i <= 8; i++ {
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

	}
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
			// fmt.Println(nr, "x ", from, "=>", to)
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
