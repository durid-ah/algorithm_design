package main

type Location struct {
	row int
	col int
}

func Loc(r int, c int) *Location { 
	return &Location{r, c}
}

type Bound struct {
	sRow int
	sCol int
	nRow int
	nCol int
}

func NewBound(sRow int, sCol int, nRow int, nCol int) *Bound {
	return &Bound {sRow, sCol, nRow, nCol}
}