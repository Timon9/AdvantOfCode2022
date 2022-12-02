package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func day2() {

	godotenv.Read(".env")

	fmt.Println("Day 2")

	fmt.Println(GetInput())

}
