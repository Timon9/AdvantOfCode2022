package main

import (
	"testing"
)

/**
0[3][0][3][7][3]
1[2][5][5][1][2]
2[6][5][3][3][2]
3[3][3][5][4][9]
4[3][5][3][9][0]
  0  1  2  3  4
LR 2
0[3][0][3][7][3]
1[2]{5}[5][1][2]
2[6][5][3][3][2]
3[3][3]{5}[4][9]
4[3][5][3][9][0]
  0  1  2  3  4

TB 1
0[3][0][3][7][3]
1[2]{5}{5}[1][2]
2[6][5][3][3][2]
3[3][3]{5}[4][9]
4[3][5][3][9][0]
  0  1  2  3  4

RL  2
0[3][0][3][7][3]
1[2]{5}{5}[1][2]
2[6]{5}[3]{3}[2]
3[3][3]{5}[4][9]
4[3][5][3][9][0]
  0  1  2  3  4

BT  0
0[3][0][3][7][3]
1[2]{5}{5}[1][2]
2[6]{5}[3]{3}[2]
3[3][3]{5}[4][9]
4[3][5][3][9][0]
  0  1  2  3  4

TEST 2

x   0 {1}{2}{4}{5}{3}
 	1 {3}{8}[4]{7}{3}
	2 {1}{3}{4}{5}{6}
	3 {9}[4][2]{4}{1}
	4 {3}{4}{1}{3}{2}
	   0  1  2  3  4 y

	   16 + lr4+tb1+rl1+bt1=
*/
func TestCanFindVisibleTrees(t *testing.T) {
	input := `30373 
25512 
65332 
33549 
35390`
	if r := findVisibleTrees(input); r != 21 {
		t.Fatalf("Part 1 FindVisibleTrees: Expected 21 but got %v", r)
	}

	input = `12453 
	38473 
	13456 
	94241 
	34132`
	if r := findVisibleTrees(input); r != 23 {
		t.Fatalf("Part 1 *2* FindVisibleTrees: Expected 23 but got %v", r)
	}
}
