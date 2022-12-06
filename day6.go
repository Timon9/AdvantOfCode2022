package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func detectMarker(input string) int {
	count := len(input)
	u := 0
	hm := make(map[byte]bool) // hashmap
	for i := 0; i < count; i++ {
		v := input[i]
		if _, ok := hm[v]; ok { // Duplicate
			for k := range hm { // Clear hash map
				delete(hm, k)
			}
			u = 0
			continue
		}

		hm[v] = true
		u++
		if u == 4 {
			return i - 2
		}

	}
	return 0
}

func day6Part1(input string) {
	fmt.Println(detectMarker(input))
}

func day6() {

	godotenv.Read(".env")

	fmt.Println("Day 6")

	day6Part1(GetInput(6))

}
