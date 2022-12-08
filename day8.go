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
		// fmt.Println("")
		// fmt.Printf("%v|", string(l))
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
		l := input[x][0]
		// fmt.Println("")
		// fmt.Printf("%v|", string(l))
		for y := len(input[x]) - 1; y > 0; y-- {
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
	fmt.Println("RL| Found", r, "int. trees")
	return r, hm
}

func tb(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for y := 1; y < len(input[0]); y++ {
		l := input[0][y]
		// fmt.Printf("%v|", string(l))

		for x := 1; x < len(input)-1; x++ {
			v := input[x][y]
			k := fmt.Sprintf("x%vy%v", y, x)
			if v > l {
				// fmt.Printf("[%v]", string(v))

				l = v
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
					// fmt.Printf("[Y]")
				}
				// else {
				// 	// fmt.Printf("[D]")

				// }
			} else {
				// fmt.Printf("%v", string(v))

			}
		}
		// fmt.Println("\n===")

	}
	fmt.Println("TB| Found", r, "int. trees ")
	return r, hm
}

func bt(input []string, hm map[string]bool) (int, map[string]bool) {
	r := 0
	for y := 1; y < len(input[0]); y++ {
		l := input[0][y]
		// fmt.Printf("%v|", string(l))

		for x := len(input) - 1; x > 1; x-- {
			v := input[x][y]
			k := fmt.Sprintf("x%vy%v", y, x)
			if v > l {
				// fmt.Printf("[%v]", string(v))

				l = v
				if _, ok := hm[k]; !ok {
					r++
					hm[k] = true
					// fmt.Printf("[Y]")
				}
				// else {
				// 	// fmt.Printf("[D]")

				// }
			} else {
				// fmt.Printf("%v", string(v))

			}
		}
		// fmt.Println("\n===")

	}
	fmt.Println("BT| Found", r, "int. trees ")
	return r, hm
}

func findVisibleTrees(input string) int {
	r := 0
	lines := strings.Split(input, "\n")
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
	fmt.Println("Result part 1", findVisibleTrees(input))
}

func day8() {

	godotenv.Read(".env")

	fmt.Println("Day 8")

	input := GetInput(8)
	day8Part1(input)

}
