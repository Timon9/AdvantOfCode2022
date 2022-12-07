package main

import (
	"testing"
)

func TestCanFindTotalSizeOfLargeDirectories(t *testing.T) {

	testData := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	totalSize := findTotalSizeOfLargeDirectories(testData)

	if totalSize != 95437 {
		t.Fatalf("Expected 95437, but got %d", totalSize)
	}
}

func TestCanFindSmallestDirectory(t *testing.T) {

	testData := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	totalSize := findSmallestDir(testData)

	if totalSize != 24933642 {
		t.Fatalf("Test part2: Expected 24933642, but got %d", totalSize)
	}
}
