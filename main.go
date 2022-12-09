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
func GetInput(day int) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	session := os.Getenv("session")

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)

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

	day9()

}
