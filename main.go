package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

/*
*
Load the input provided by AdventOfCode.

@return string
*
*/
func GetInput() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	session := os.Getenv("session")

	url := "https://adventofcode.com/2022/day/2/input"

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func main() {

	godotenv.Read(".env")

	fmt.Println("AdvantOfCode")

	day1()

}
