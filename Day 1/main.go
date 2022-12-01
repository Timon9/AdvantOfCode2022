package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

/**
Load the input provided by AdventOfCode.

@return string
**/
func getInput() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	session := os.Getenv("session")

	url := "https://adventofcode.com/2022/day/1/input"

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func main() {

	godotenv.Read(".env")

	fmt.Println("Day 1")

	input := strings.Split(getInput(), "\n") // Break the string into an array

	/**
	Solve the puzzle in O(n)
	**/
	c := 0 // Highest sub-count
	d := 0 // Second highest sub-count
	e := 0 // Third highest sub-count

	s := 0 // Counter per Elv
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			s = 0 // Reset the counter, an empty line indicates a new Elv
		} else {
			p, _ := strconv.Atoi(input[i]) // Convert string to int
			s = s + p
			if s > c { // Highest counter
				c = s
			} else if s > d { // Second highest counter
				d = s
			} else if s > e { // Third highest counter
				e = s
			}

		}
	}

	fmt.Println("Solution is:", (c + d + e))

}
