package main

import (
	gss "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func main() {
	b := gss.Board{}
	b.TakeInput()
	b.Solve()
}
