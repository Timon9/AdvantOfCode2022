package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func findTotalSizeOfLargeDirectories(input string) int {

	lines := strings.Split(input, "\n")
	cm := make(map[string]int)
	rm := []string{"/"}

	dl := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		ll := len(line)
		fmt.Println("==")

		if ll > 1 && line[0:1] == "$" {
			dl = false
		}

		if ll > 5 && line[0:5] == "$ cd " {
			nextDir := line[5:ll]
			switch nextDir {
			case "/":
			case "..":
				fmt.Println("Switching")
				fmt.Println(rm)
				rm = rm[0 : len(rm)-1]
				fmt.Println(rm)
				continue
			default:
				rm = append(rm, nextDir)
			}
		}

		if dl == true && len(line) > 3 && line[0:3] != "dir" {
			tmp := strings.Split(line, " ")
			size, _ := strconv.Atoi(tmp[0])
			for ii := 0; ii < len(rm); ii++ {
				cm[rm[ii]] = cm[rm[ii]] + size
			}
		}

		if ll > 3 && line[0:4] == "$ ls" {
			dl = true
		}

	}

	r := 0

	// fmt.Println(cm)
	for k := range cm {
		if cm[k] < 100000 {
			r = r + cm[k]
		}
	}

	return r
}

func day7Part1(input string) {
	//1104798<
	fmt.Println("Result", findTotalSizeOfLargeDirectories(input))
}

func day7() {

	godotenv.Read(".env")

	fmt.Println("Day 7")

	input := GetInput(7)
	day7Part1(input)

}
