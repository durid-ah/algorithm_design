package main


// PeakProblem A class representing an instance of a peak-finding problem.
type PeakProblem struct {
	array    [][]int
	startRow int
	startCol int
	numRow   int
	numCol   int
}


func (peak *PeakProblem) Get(location *Location) int {
	r := location.row
	c := location.col

	if !(0 <= r && r < peak.numRow) {
		return 0
	} else if !(0 <= c && c < peak.numCol) {
		return 0
	}

	return peak.array[peak.startRow+r][peak.startCol+c]
}


func (peak *PeakProblem) getBetterNeighbor(location *Location) *Location {
	r, c := location.row, location.col
	best := location

	tempLoc := Location{r - 1, c}
	if r-1 >= 0 && peak.Get(&tempLoc) > peak.Get(best) {
		best = &tempLoc
	}

	tempLoc = Location{r, c - 1}
	if c-1 >= 0 && peak.Get(&tempLoc) > peak.Get(best) {
		best = &tempLoc
	}

	tempLoc = Location{r + 1, c}
	if r+1 < peak.numRow && peak.Get(&tempLoc) > peak.Get(best) {
		best = &tempLoc
	}

	tempLoc = Location{r, c + 1}
	if c+1 < peak.numCol && peak.Get(&tempLoc) > peak.Get(best) {
		best = &tempLoc
	}

	return best
}


func (peak *PeakProblem) getMaximum(locations []Location) *Location {
	var bestLoc *Location = nil
	bestVal := 0

	for _, loc := range locations {
		locVal := peak.Get(&loc)
		if bestLoc == nil || locVal > bestVal {
			bestLoc = &loc
			bestVal = locVal
		}
	}

	return bestLoc
}

func (peak *PeakProblem) isPeak(location *Location) bool {
	betterLoc := peak.getBetterNeighbor(location)
	return betterLoc == location
}

func (peak *PeakProblem) getSubproblem(sRow int, sCol int, nRow int, nCol int) *PeakProblem {
	newPeak := PeakProblem{
		peak.array,
		peak.startRow + sRow,
		peak.startCol + sCol,
		nRow, nCol}

	return &newPeak
}

func (peak *PeakProblem) getSubproblemContaining(
	boundList []Bound, location *Location) *PeakProblem {

	row := location.row
	col := location.col

	for _, bound := range boundList {
		if (bound.sRow <= row) &&
			(row < (bound.sRow + bound.nRow)) {
			if (bound.sCol <= col) &&
				(col < bound.sCol+bound.nCol) {
				return peak.getSubproblem(
					bound.sRow, bound.sCol, bound.nRow, bound.nCol)
			}
		}
	}

	return peak
}

// Remaps the location in the given problem to the same location in
// the problem that this function is being called from.
//
// RUNTIME: O(1)
func (peak *PeakProblem)getLocationInSelf(
	problem *PeakProblem, location *Location) Location {

	row := location.row
	col := location.col

	newRow := row + problem.startRow - peak.startRow
	newCol := col + problem.startCol - peak.startCol

	return Location{newRow, newCol}
}

/***********************************************************************/
/* Helpers																				  */
/***********************************************************************/

// Gets the dimensions for a two-dimensional array.  The first dimension
// is simply the number of items in the list; the second dimension is the
// length of the shortest row.  This ensures that any location (row, col)
// that is less than the resulting bounds will in fact map to a valid
// location in the array.
//
// RUNTIME: O(len(array))
func getDimensions(array *[][]int) (int, int) {
	rows := len(*array)
	cols := 0

	for _, row := range *array {
		if len(row) > cols {
			cols = len(row)
		}
	}

	return rows, cols
}

// Constructs an instance of the PeakProblem object for the given array,
// using bounds derived from the array using the getDimensions function.
//
// RUNTIME: O(len(array))
func createProblem(array [][]int) *PeakProblem {
	rows, cols := getDimensions(&array)
	return &PeakProblem{array, 0, 0, rows, cols}
}