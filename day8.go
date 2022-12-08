package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func lr(input []string) (int, []string) {
	r := 0
	for i := 1; i < len(input)-1; i++ {
		v := input[i]
		l := v[0]
		fmt.Println("")
		fmt.Printf("%v|", string(l))
		for ii := 0; ii < len(v)-1; ii++ {
			if v[ii] > l {
				l = v[ii]
				r++
				fmt.Printf("[%v]", string(v[ii]))
				input[i] = input[i][0:ii] + "0" + input[i][ii+1:len(input[i])]

			} else {
				fmt.Printf("%v.", string(v[ii]))
			}

		}

	}
	fmt.Println("Found", r, "int. trees lr")
	return r, input
}

func findVisibleTrees(input string) int {
	r := 0
	lines := strings.Split(input, "\n")
	c := len(lines)
	a := 0
	for i := 0; i < len(lines[0]); i++ {
		a++
	}

	out := c + (a - 1) + (a - 1) + (c - 2)

	r = out
	x, y := lr(lines)
	r = r + x
	lr(y)
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
