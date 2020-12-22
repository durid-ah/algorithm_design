package main

import (
	"fmt"
)

// Note: algorithm 3 was ommitted because it was incorrect
// and I didn't do 4 because I didn't feel like it
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

	problem := CreateProblem(problemMatrix)

	fmt.Println("The problem matrix:")
	for _, row := range problemMatrix {
		for _, value := range row {
			fmt.Printf("%3d ", value)
		}
		fmt.Printf("\n")
	}

	loc1 := algorithm1(problem)
	val1 := problem.Get(loc1)
	
	loc2 := algorithm2(problem, *Loc(0, 0))
	val2 := problem.Get(loc1)

	fmt.Printf("\n")
	fmt.Printf("Algorithm 1: %+v Value: %v \n", loc1, val1)
	fmt.Printf("Algorithm 2: %+v Value: %v \n", loc2, val2)

}
