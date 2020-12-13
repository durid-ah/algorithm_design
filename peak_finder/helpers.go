package main


// Location of a value (i.e Peak)
type Location struct {
	row int
	col int
}


// Loc Instatiantes a Location struct
func Loc(r int, c int) *Location {
	return &Location{r, c}
}


// Bound contains the bounds of the exercise
type Bound struct {
	sRow int
	sCol int
	nRow int
	nCol int
}

// IntPair a pair of two int types
type IntPair struct {
	first 	int
	second 	int
}


// NewBound Instantiates a Bound struct
func NewBound(sRow int, sCol int, nRow int, nCol int) *Bound {
	return &Bound{sRow, sCol, nRow, nCol}
}

// CrossProduct Returns all pairs with one item from the first list and one item from 
// the second list.  (Cartesian product of the two lists.)
func CrossProduct(list1 []int, list2 []int) []Location {
	var answer []Location
	for _, a := range list1 {
		for _, b := range list2 {
			answer = append(answer, Location{a, b})
		}
	}

	return answer
}

// MakeRange creates a list of consecutive numbers
func MakeRange(min int, max int) []int {
	a := make([]int, max - min)
	for i := range a {
		a[i] = min + i
  }
  return a
}