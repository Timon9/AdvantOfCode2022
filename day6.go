package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func detectMarker(input string, distinct int) int {
	count := len(input)
	hm := make(map[int]map[byte]bool) // hashmap
	for i := distinct; i < count; i++ {
		hm[i] = make(map[byte]bool)
		s := i - distinct
		if s < 0 {
			s = 0
		}
		v := input[s:i]
		for ii := 0; ii < len(v); ii++ {
			vv := v[ii]
			if _, ok := hm[i][vv]; !ok { // Unique
				hm[i][vv] = true
				if len(hm[i]) >= distinct {
					return i
				}
			}

		}

	}
	return 0
}

func day6Part1(input string) {
	fmt.Println("Day6Part1", detectMarker(input, 4))
}
func day6Part2(input string) {
	fmt.Println("Day6Part2", detectMarker(input, 14))
}

func day6() {

	godotenv.Read(".env")

	fmt.Println("Day 6")

	input := GetInput(6)
	day6Part1(input)
	day6Part2(input)

}
