package main

// algorithm2 starts at a specified location and moves towards a higher
// value
func algorithm2(problem *PeakProblem, location Location) *Location {
	if problem.numRow <= 0 || problem.numCol <= 0 {
		return nil
	}

	nextLoc := problem.GetBetterNeighbor(&location)

	if *nextLoc == location {
		return &location
	} else {
		return algorithm2(problem, *nextLoc)
	}
}
