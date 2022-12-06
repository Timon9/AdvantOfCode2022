package main

import (
	"testing"
)

func TestCanDetectMarker(t *testing.T) {
	firstTest := detectMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if firstTest != 5 {
		t.Fatalf("First test expected 5, got %d", firstTest)
	}
	firstTest := detectMarker("nppdvjthqldpwncqszvftbrmjlhg")
	if firstTest != 6 {
		t.Fatalf("First test expected 6, got %d", firstTest)
	}
	firstTest := detectMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if firstTest != 10 {
		t.Fatalf("First test expected 10, got %d", firstTest)
	}
	firstTest := detectMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	if firstTest != 11 {
		t.Fatalf("First test expected 11, got %d", firstTest)
	}

}
