package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func buildDirTree(input string) map[string]int {
	lines := strings.Split(input, "\n")
	cm := make(map[string]int)
	rm := []string{}

	dl := false

	ttmp := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		ll := len(line)

		if ll > 1 && line[0:1] == "$" {
			dl = false
		}

		if ll > 5 && line[0:5] == "$ cd " {
			ttmp++
			switch line[5:ll] {
			case "..":
				rm = rm[0 : len(rm)-1]
				continue
			default:
				nextDir := fmt.Sprintf("lv_%d_%v", ttmp, line[5:ll])
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
	return cm
}
func findSmallestDir(input string) int {

	cm := buildDirTree(input)
	r := 30000000

	totalSize := cm["lv_1_/"]
	unused := 70000000 - totalSize
	needed := 30000000 - unused

	for k := range cm {
		if cm[k] >= needed && cm[k] < r {
			r = cm[k]
		}
	}

	return r
}

func findTotalSizeOfLargeDirectories(input string) int {

	cm := buildDirTree(input)

	r := 0

	for k := range cm {
		if cm[k] < 100000 {
			r = r + cm[k]
		}
	}

	return r
}

func day7Part1(input string) {
	fmt.Println("Result part 1", findTotalSizeOfLargeDirectories(input))
}

func day7Part2(input string) {
	fmt.Println("Result part 2", findSmallestDir(input))
}

func day7() {

	godotenv.Read(".env")

	fmt.Println("Day 7")

	input := GetInput(7)
	day7Part1(input)
	day7Part2(input)

}
