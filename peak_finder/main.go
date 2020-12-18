package main

import (
	"fmt"
)

func main() {
	problemMatrix := [][]int{
		{4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2},
		{5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3},
		{6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4},
		{7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5},
		{8, 9, 10, 11, 12, 11, 10, 9, 8, 7, 6},
		{7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5},
		{6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4},
		{5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3},
		{4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2},
		{3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 5, 4, 3, 2, 1, 0},
	}

	fmt.Println(problemMatrix)

	//fmt.Println(MakeRange(10, 20))
	fmt.Println(algorithm1(CreateProblem(problemMatrix)))
	fmt.Println(algorithm2(CreateProblem(problemMatrix), *Loc(0,0)))
}
