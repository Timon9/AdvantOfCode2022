package main

import (
	"fmt"
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
	fmt.Println("LR| Found", r, "int. trees")
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
			fmt.Println(k)
			if input[x][y] > l {
				l = input[x][y]
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
				}

			}

		}

	}
	fmt.Println("RL| Found", r, "int. trees")
	return r, hm
}

func tb(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for y := 1; y < len(input[0])-1; y++ {
		l := input[0][y]
		for x := 1; x < len(input)-1; x++ {
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
	fmt.Println("TB| Found", r, "int. trees ")
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
	fmt.Println("BT| Found", r, "int. trees ")
	return r, hm
}

func findVisibleTrees(input string) int {
	r := 0
	lines := strings.Split(input, "\n")
	for a := 0; a < len(lines); a++ {
		lines[a] = strings.TrimSpace(lines[a])
	}
	hm := make(map[string]bool)
	c := len(lines)
	a := 0
	for i := 0; i < len(lines[0]); i++ {
		a++
	}

	out := c + (a - 1) + (a - 1) + (c - 2)

	r = out
	x, y := lr(lines, hm)
	xx, yy := tb(lines, y)
	xxx, yyy := rl(lines, yy)
	xxxx, _ := bt(lines, yyy)
	r = r + x + xx + xxx + xxxx

	return r
}

func day8Part1(input string) {
	//<1773
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day8() {

	godotenv.Read(".env")

	fmt.Println("Day 8")

	input := GetInput(8)
	day8Part1(input)

}
