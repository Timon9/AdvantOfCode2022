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

	input := strings.Split(getInput(), "\n")

	/**
	Solve the puzzle in O(n)
	**/
	c := 0
	d := 0
	e := 0

	s := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			s = 0
		} else {
			p, _ := strconv.Atoi(input[i])
			s = s + p
			if s > c {
				c = s
			} else if s > d {
				d = s
			} else if s > e {
				e = s
			}

		}
	}

	fmt.Println("Solution is:", (c + d + e))

}
