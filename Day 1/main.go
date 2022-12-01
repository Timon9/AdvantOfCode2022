package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getInput() string {
	url := "https://adventofcode.com/2022/day/1/input"
	session := "53616c7465645f5f8e63855ae159eac1a344ad8f62990711f5d66329491c39a836d60350a7a51ca9f8a0907b96eb2c83766ad6c93dfafcbc87dbdab4bcb75872"

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func main() {
	fmt.Println("Day 1")

	input := getInput()

	fmt.Println(input)

}
