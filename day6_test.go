package main

import (
	"testing"
)

func TestCanDetectMarker(t *testing.T) {
	firstTest := detectMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4) //b|vwbj|plbgvbhsrlpgdmjqwftvncz
	if firstTest != 5 {
		t.Fatalf("First test expected 5, got %d", firstTest)
	}
	secondTest := detectMarker("nppdvjthqldpwncqszvftbrmjlhg", 4) //np|pdvj|thqldpwncqszvftbrmjlhg
	if secondTest != 6 {
		t.Fatalf("secondTest test expected 6, got %d", firstTest)
	}
	thirdTest := detectMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4) //nznrnf|rfnt|jfmvfwmzdfjlvtqnbhcprsg
	if thirdTest != 10 {
		t.Fatalf("thirdTest test expected 10, got %d", firstTest)
	}
	fourtTest := detectMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	if fourtTest != 11 {
		t.Fatalf("fourtTest test expected 11, got %d", firstTest)
	}

}

func TestCanDetectMessage(t *testing.T) {
	firstTest := detectMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14) //mjqjp|qmgbljsphdztnv|jfqwrcgsmlb
	if firstTest != 19 {
		t.Fatalf("First message test expected 19, got %d", firstTest)
	}
	secondTest := detectMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14) //mjqjp|qmgbljsphdztnv|jfqwrcgsmlb
	if secondTest != 23 {
		t.Fatalf("Second message test expected 23, got %d", firstTest)
	}

	thirdTest := detectMarker("nppdvjthqldpwncqszvftbrmjlhg", 14) //mjqjp|qmgbljsphdztnv|jfqwrcgsmlb
	if thirdTest != 23 {
		t.Fatalf("Third message test expected 23, got %d", firstTest)
	}
	fourtTest := detectMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14) //mjqjp|qmgbljsphdztnv|jfqwrcgsmlb
	if fourtTest != 29 {
		t.Fatalf("Third message test expected 29, got %d", firstTest)
	}

}
