package main

import (
	"fmt"
)

func algorithm1(problem *PeakProblem) *Location {
	// If it's empty, we're done
	if problem.numCol <= 0 && problem.numRow <= 0 {
		return nil
	}

	// the recursive subproblem will involve half the number of columns
	// find the middle column through integer division
	mid := problem.numCol / 2

	// Building the two subproblems
	subStartR, subNumR := 0, problem.numRow //Both subproblems will have the same number of rows

	// subproblem_1 will start with index 0
	// mid is used as the number of collumns in subproblem_1
	subStartC1, subNumC1 := 0, mid

	// second subproblem will start from the middle + 1
	// subNumC2 is the number of columns of the subproblem
	subStartC2, subNumC2 := mid+1, problem.numCol-(mid+1)

	var subproblems []Bound
	subproblems = append(subproblems, Bound{subStartR, subStartC1, subNumR, subNumC1})
	subproblems = append(subproblems, Bound{subStartR, subStartC2, subNumR, subNumC2})

	divider := CrossProduct(MakeRange(0, problem.numRow), []int{mid})
	fmt.Println(divider)

	return nil
}
