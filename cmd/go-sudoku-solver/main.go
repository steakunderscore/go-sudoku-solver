package main

import (
	"os"

	gss "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func main() {
	b := gss.Board{}
	b.TakeInput()
	b.Solve()
	b.Print(os.Stdout)
}
