package main

type Location struct {
	row int
	col int
}

func Loc(r int, c int) *Location { 
	return &Location{r, c}
}