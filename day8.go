package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func lr(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for x := 1; x < len(input)-1; x++ {
		l := input[x][0]
		for y := 0; y < len(input[x])-1; y++ {
			k := fmt.Sprintf("x%vy%v", x, y)
			if input[x][y] > l {
				l = input[x][y]
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
				}

			}

		}

	}
	return r, hm
}

func rl(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for x := 1; x < len(input)-1; x++ {
		if len(input[x]) <= 1 {
			continue
		}
		l := input[x][len(input[x])-1]
		for y := len(input[x]) - 2; y > 0; y-- {
			k := fmt.Sprintf("x%vy%v", x, y)
			if input[x][y] > l {
				l = input[x][y]
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
				}

			}

		}

	}
	return r, hm
}

func tb(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for y := 1; y < len(input[0])-1; y++ {
		l, _ := strconv.Atoi(string(input[0][y]))
		for x := 1; x < len(input)-1; x++ {
			v, _ := strconv.Atoi(string(input[x][y]))
			k := fmt.Sprintf("x%vy%v", x, y)
			if v > l {
				l = v
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
				}
			}
		}

	}
	return r, hm
}

func bt(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for y := 1; y < len(input[0])-1; y++ {
		l := input[len(input[0])-1][y]

		for x := len(input) - 2; x > 1; x-- {
			v := input[x][y]
			k := fmt.Sprintf("x%vy%v", x, y)
			if v > l {

				l = v
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
				}
			}
		}

	}
	return r, hm
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func findVisibleTrees(input string) int {

	r := 0
	lines := strings.Split(input, "\n")
	for a := 0; a < len(lines); a++ {
		lines[a] = strings.TrimSpace(lines[a])
		if len(lines[a]) < 10 {
			lines = removeIndex(lines, a)
		}
	}

	hm := make(map[string]bool)
	c := len(lines)
	out := ((c * 2) + (len(lines[0]) * 2)) - 4

	lr, y := lr(lines, hm)
	tb, yy := tb(lines, y)
	rl, yyy := rl(lines, yy)
	bt, _ := bt(lines, yyy)

	r = out + lr + tb + rl + bt

	return r
}

func findUp(lines []string, x int, y int) int {
	s, _ := strconv.Atoi(string(lines[x][y]))
	c := 0
	for i := x - 1; i >= 0; i-- {
		c++
		b, _ := strconv.Atoi(string(lines[i][y]))
		if b >= s {
			return c
		}

	}
	return c
}

func findDown(lines []string, x int, y int) int {
	s, _ := strconv.Atoi(string(lines[x][y]))
	c := 0
	for i := x + 1; i < len(lines); i++ {
		c++
		b, _ := strconv.Atoi(string(lines[i][y]))
		if b >= s {
			return c
		}

	}
	return c
}

func findLeft(lines []string, x int, y int) int {
	s, _ := strconv.Atoi(string(lines[x][y]))
	c := 0
	for ii := y - 1; ii >= 0; ii-- {
		c++
		b, _ := strconv.Atoi(string(lines[x][ii]))

		if b >= s {
			return c
		}

	}
	return c
}

func findRight(lines []string, x int, y int) int {
	s, _ := strconv.Atoi(string(lines[x][y]))
	c := 0
	for ii := y + 1; ii < len(lines[x]); ii++ {
		c++
		b, _ := strconv.Atoi(string(lines[x][ii]))

		if b >= s {
			return c
		}

	}
	return c
}

func findBestTree(input string) int {

	r := 0
	lines := strings.Split(input, "\n")
	for a := 0; a < len(lines); a++ {
		lines[a] = strings.TrimSpace(lines[a])
		if len(lines[a]) < 1 {
			lines = removeIndex(lines, a)
		}
	}

	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			up := findUp(lines, x, y)
			down := findDown(lines, x, y)
			left := findLeft(lines, x, y)
			right := findRight(lines, x, y)

			tree := up * left * down * right
			if tree > r {
				r = tree
			}
		}

	}

	return r
}

func day8Part1(input string) {
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day8Part2(input string) {
	fmt.Println("Result part2", findBestTree(input))
}

func day8() {

	godotenv.Read(".env")

	fmt.Println("Day 8")

	input := GetInput(8)
	day8Part2(input)

}
