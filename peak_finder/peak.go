package main

type PeakProblem struct {
	array [][]int
	startRow int
	startCol int
	numRow int
	numCol int
}

func (self * PeakProblem) Get(r int, c int) int {

	if !(0 <= r && r < self.numRow) {
		return 0
	} else if !(0 <= c && c < self.numCol) {
		return 0
	} 
		
	return self.array[self.startRow + r][self.startCol + c]
}

func (self *PeakProblem)getBetterNeighbor(r int, c int) (int, int) {
	bestr := r
	bestc := c
	
	if r - 1 >= 0 && self.Get(r - 1, c) > self.Get(bestr, bestc) {
		bestr = r - 1
		bestc = c
	} 
	
	if c - 1 >= 0 && self.Get(r, c - 1) > self.Get(bestr, bestc) {
		bestr = r
		bestc = c - 1		
	}
	
	if r + 1 < self.numRow && self.Get(r + 1, c) > self.Get(bestr, bestc) {
		bestr = r + 1
		bestc = c
	}

	if c + 1 < self.numCol && self.Get(r, c + 1) > self.Get(bestr, bestc) {
		bestr = r + 1
		bestc = c
	}

	return bestr, bestc
}


