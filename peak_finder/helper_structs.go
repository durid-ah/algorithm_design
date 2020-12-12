package main


type Location struct {
	row int
	col int
}


// Loc Instatiantes a Location struct
func Loc(r int, c int) *Location {
	return &Location{r, c}
}

type Bound struct {
	sRow int
	sCol int
	nRow int
	nCol int
}

// NewBound Instantiates a Bound struct
func NewBound(sRow int, sCol int, nRow int, nCol int) *Bound {
	return &Bound{sRow, sCol, nRow, nCol}
}
