package main

type PeakProblem struct {
	array [][]int
	startRow int
	startCol int
	numRow int
	numCol int
}

func (self * PeakProblem) Get(location *Location) int {
	r := location.row
	c := location.col

	if !(0 <= r && r < self.numRow) {
		return 0
	} else if !(0 <= c && c < self.numCol) {
		return 0
	} 
		
	return self.array[self.startRow + r][self.startCol + c]
}

func (self *PeakProblem)getBetterNeighbor(location *Location) *Location {
	r, c := location.row, location.col
	best := location
	
	tempLoc := Location{r - 1, c}
	if r - 1 >= 0 && self.Get(&tempLoc) > self.Get(best) {
		best = &tempLoc
	} 
	
	tempLoc = Location{r, c - 1}
	if c - 1 >= 0 && self.Get(&tempLoc) > self.Get(best) {
		best = &tempLoc
	}
	
	tempLoc = Location{r + 1, c}
	if r + 1 < self.numRow && self.Get(&tempLoc) > self.Get(best) {
		best = &tempLoc
	}

	tempLoc = Location{r , c + 1}
	if c + 1 < self.numCol && self.Get(&tempLoc) > self.Get(best) {
		best = &tempLoc
	}

	return best
}

func (self *PeakProblem) getMaximum(locations []Location) *Location {
	var bestLoc *Location = nil
	bestVal := 0

	for _, loc := range locations {
		locVal := self.Get(&loc)
		if bestLoc == nil || locVal > bestVal {
			bestLoc = &loc
			bestVal = locVal
		}
	}

	return bestLoc
}

func (self *PeakProblem) isPeak(location *Location) bool {
	betterLoc := self.getBetterNeighbor(location)
	return betterLoc == location
}

func (peak *PeakProblem) getSubproblem(sRow int, sCol int, nRow int, nCol int) *PeakProblem {
	newPeak := PeakProblem { 
		peak.array,
		peak.startRow + sRow,
		peak.startCol + sCol,
		nRow, nCol }
	
	return &newPeak
}

func (self *PeakProblem) getSubproblemContaining(boundList []Bound) *Location {
 	return nil 
}


